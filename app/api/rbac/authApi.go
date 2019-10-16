package rbac

import (
	"github.com/gogf/gf/net/ghttp"
	. "github.com/zhwei820/gadmin/app/api/base"
)

//
// @Summary 登出 （jwt 后端暂时无法登出）
// @Description 登出
// @Tags auth
// @Success 200 {string} string	"ok"
// @router /rbac/logout [get]
func Logout(r *ghttp.Request) {
	ReqDebug(r, "ReqDebug ReqDebug: %v", struct {
		a string
	}{"11235478"})
	Success(r, "success")
	ReqInfo(r, "ReqInfo ReqInfo %v", []int{1, 2, 9000})

}

//
// @Summary 登陆
// @Description 登陆
// @Tags auth
// @Param   Login  body api_model.Login true "Login"
// @Success 200 {string} string	"ok"
// @router /rbac/login [post]
func SwaggerLogin(r *ghttp.Request) {
	Success(r, "success")
}

//
// @Summary 刷新token
// @Description 刷新token
// @Tags auth
// @Success 200 {string} string	"ok"
// @router /rbac/refresh_token [get]
func SwaggerRefreshToken(r *ghttp.Request) {
	Success(r, "success")
}
