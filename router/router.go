package router

import (
	"fmt"
	"github.com/hailaz/gadmin/app/service"
	"strings"

	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
	"github.com/hailaz/gadmin/app/api"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/common"
)

var routerMap = make(map[string]model.RolePolicy)

func showURL(r *ghttp.Request) {
	glog.Debug("请求路径：", r.Method, r.Request.RequestURI)
}

// InitRouter 初始化路由
//
// createTime:2019年05月13日 09:32:58
// author:hailaz
func InitRouter(s *ghttp.Server) {

	s.BindHookHandler("/*", ghttp.HOOK_BEFORE_SERVE, showURL)
	Init(s)

	service.ReSetPolicy("system", routerMap)
}

// authHook 鉴权钩子
//
// createTime:2019年05月13日 09:33:58
// author:hailaz
// authHook is the HOOK function implements JWT logistics.
func authHook(r *ghttp.Request) {
	uri := strings.Split(r.Request.RequestURI, "/")
	if len(uri) > 1 {
		switch uri[1] { //登录相关免鉴权
		case "loginkey":
			return
		case "login":
			return
		case "swagger":
			return
		}
	}

	r.Response.CORSDefault() //开启跨域
	//r.Response.Header().Set("Access-Control-Allow-Origin", "*")
	api.GfJWTMiddleware.MiddlewareFunc()(r) //鉴权中间件
}

type Object struct{}

func (o *Object) Show(r *ghttp.Request) {
	r.Response.Writeln("Object Show")
}

func (o *Object) Delete(r *ghttp.Request) {
	r.Response.Writeln("Object REST Delete")
}

// Init 初始化V1
//
// createTime:2019年04月25日 09:24:06
// author:hailaz
func Init(s *ghttp.Server) {
	//权限验证
	s.Group("").ALL("/*any", authHook, ghttp.HOOK_BEFORE_SERVE)
	userCtrl := new(api.UserController)
	roleCtrl := new(api.RoleController)
	policyCtrl := new(api.PolicyController)
	menuCtrl := new(api.MenuController)

	// user
	BindGroup(s, "/", []ghttp.GroupItem{
		//
		//登录
		{"GET", "/loginkey", api.GetLoginCryptoKey},                   //获取登录加密公钥
		{"POST", "/login", api.GfJWTMiddleware.LoginHandler},          //登录
		{"GET", "/refresh_token", api.GfJWTMiddleware.RefreshHandler}, //获取登录加密公钥
		{"POST", "/logout", api.Logout},                               //登出
		////menu
		{"REST", "/menu", menuCtrl},
		//// 用户
		{"GET", "/user/info", userCtrl, "Info", "false"},
		{"GET", "/user/menu", userCtrl, "Menu", "false"},
		{"REST", "/user", userCtrl},
		// 角色
		{"REST", "/role", roleCtrl},
		{"PUT", "/role/byuser", roleCtrl, "SetRoleByUserName"},
		{"PUT", "/role/menu", roleCtrl, "SetRoleMenus"},
		// 权限
		{"REST", "/policy", policyCtrl},
		{"GET", "/policy/byrole", policyCtrl, "GetPolicyByRole"},
		{"PUT", "/policy/byrole", policyCtrl, "SetPolicyByRole"},
	})
}

// BindGroup 绑定分组路由
//
// createTime:2019年04月29日 16:45:55
// author:hailaz
func BindGroup(s *ghttp.Server, path string, items []ghttp.GroupItem) {
	g := s.Group(path)
	g.Bind(items)
	for _, item := range items {
		glog.Debug(gconv.String(item[1]))
		if len(item) > 4 && gconv.String(item[4]) == "false" { //不走权限的api
			addPolicy("*", path+gconv.String(item[1]), common.GetAction(gconv.String(item[0])))
		} else { //走权限的api
			if gconv.String(item[0]) == "REST" { //rest api
				addPolicy("system", path+gconv.String(item[1]), model.ACTION_GET)
				addPolicy("system", path+gconv.String(item[1]), model.ACTION_POST)
				addPolicy("system", path+gconv.String(item[1]), model.ACTION_PUT)
				addPolicy("system", path+gconv.String(item[1]), model.ACTION_DELETE)
			} else {
				addPolicy("system", path+gconv.String(item[1]), common.GetAction(gconv.String(item[0])))
			}
		}

	}

}

// addPolicy 记录需要系统路由
//
// createTime:2019年04月29日 17:18:25
// author:hailaz
func addPolicy(role, path, atc string) {
	routerMap[fmt.Sprintf("%v %v %v", role, path, atc)] = model.RolePolicy{Role: role, Path: path, Atc: atc}
}
