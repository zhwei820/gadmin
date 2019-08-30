package router

import (
	"fmt"
	"github.com/gogf/gf/g"
	"github.com/hailaz/gadmin/app/service"
	"github.com/hailaz/gadmin/utils"
	"github.com/hailaz/gadmin/utils/context_log"
	uuid "github.com/satori/go.uuid"
	"io"
	"strings"

	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
	"github.com/hailaz/gadmin/app/api"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/common"
)

var routerMap = make(map[string]model.RolePolicy)
var writer io.Writer

func writeLog(r *ghttp.Request) {
	r.GetParam("ctx").Val().(*context_log.Context).WriteLog()
}

// InitRouter 初始化路由
//
// createTime:2019年05月13日 09:32:58
// author:hailaz
func InitRouter(s *ghttp.Server) {
	writer, _ = context_log.NewFileWriter("log/log", 1024*1024*50, 30)

	s.BindHookHandlerByMap("/*", map[string]ghttp.HandlerFunc{
		ghttp.HOOK_BEFORE_SERVE: authHook,
		ghttp.HOOK_AFTER_SERVE:  writeLog,
	})

	Init(s)

	service.ReSetPolicy("system", routerMap)
}

// authHook 鉴权钩子
//
// createTime:2019年05月13日 09:33:58
// author:hailaz
// authHook is the HOOK function implements JWT logistics.
func authHook(r *ghttp.Request) {
	uid := uuid.NewV4()
	r.SetParam("req", uid)
	r.SetParam("ctx", context_log.NewContext(r.Context(), writer, uid.String(), utils.GetLogLevel(g.Config().GetString("ReqLogLevel"))))

	uri := strings.Split(r.Request.RequestURI, "/")
	if len(uri) >= 3 {
		switch uri[1] + "/" + uri[2] { //登录相关免鉴权
		case "rbac/loginkey", "rbac/login", "rbac/logout":
			return
		}
	}
	if len(uri) >= 2 {
		switch uri[1] { //登录相关免鉴权
		case "swagger", "favicon.ico":
			return
		}
	}
	r.Response.CORSDefault() //开启跨域
	//r.Response.Header().Set("Access-Control-Allow-Origin", "*")
	api.GfJWTMiddleware.MiddlewareFunc()(r) //鉴权中间件
	// or error handling
}

// Init 初始化V1
//
// createTime:2019年04月25日 09:24:06
// author:hailaz
func Init(s *ghttp.Server) {
	userCtrl := new(api.UserController)
	roleCtrl := new(api.RoleController)
	policyCtrl := new(api.PolicyController)
	menuCtrl := new(api.MenuController)

	// user
	BindGroup(s, "/rbac", []ghttp.GroupItem{
		//
		//登录
		{"GET", "/loginkey", api.GetLoginCryptoKey},                   //获取登录加密公钥
		{"POST", "/login", api.GfJWTMiddleware.LoginHandler},          //登录
		{"GET", "/refresh_token", api.GfJWTMiddleware.RefreshHandler}, //获取登录加密公钥
		{"GET", "/logout", api.Logout},                                //登出
		////menu
		{"REST", "/menu", menuCtrl},
		//// 用户
		{"GET", "/user/info", userCtrl, "Info"},
		{"GET", "/user/menu", userCtrl, "Menu"},
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

		if gconv.String(item[0]) == "REST" { //rest api
			addPolicy("system", gconv.String(item[1]), model.ACTION_GET)
			addPolicy("system", gconv.String(item[1]), model.ACTION_POST)
			addPolicy("system", gconv.String(item[1]), model.ACTION_PUT)
			addPolicy("system", gconv.String(item[1]), model.ACTION_DELETE)
		} else {
			addPolicy("system", gconv.String(item[1]), common.GetAction(gconv.String(item[0])))
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
