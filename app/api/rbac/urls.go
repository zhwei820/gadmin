package rbac

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/zhwei820/gadmin/app/api/base"
	"github.com/zhwei820/gadmin/common"
)

// Init 初始化rbac
//
// createTime:2019年04月25日 09:24:06
// author:hailaz
func InitRouter(s *ghttp.Server) {
	userCtrl := new(UserController)
	roleCtrl := new(RoleController)
	policyCtrl := new(PolicyController)

	// user
	common.BindGroup(s, "/rbac", []ghttp.GroupItem{
		//
		//登录
		{"POST", "/login", base.GfJWTMiddleware.LoginHandler},          //登录
		{"GET", "/refresh_token", base.GfJWTMiddleware.RefreshHandler}, //获取登录加密公钥
		{"GET", "/logout", Logout},                                     //登出
		//// 用户
		{"GET", "/user/info", userCtrl, "Info"},
		{"REST", "/user", userCtrl},
		{"PUT", "/user/userrole", userCtrl, "SetUserRole"},

		// 角色
		{"REST", "/role", roleCtrl},
		{"GET", "/role/byrolekey", roleCtrl, "GetByRoleKey"},

		// 权限
		{"REST", "/policy", policyCtrl},
		//{"GET", "/policy/byrole", policyCtrl, "GetPolicyByRole"},
		//{"PUT", "/policy/byrole", policyCtrl, "SetPolicyByRole"},
	})
}
