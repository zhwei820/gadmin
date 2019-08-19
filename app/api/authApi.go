package api

import (
	"time"

	"github.com/gogf/gf-jwt"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/library/common"
)

var (
	// The underlying JWT middleware.
	GfJWTMiddleware *jwt.GfJWTMiddleware
)

// Initialization function,
// rewrite this function to customized your own JWT settings.
func init() {
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:                 "gf admin",
		Key:                   []byte("secret key"),
		Timeout:               time.Hour * 24 * 100,    //token有效时间1
		MaxRefresh:            time.Hour * 24 * 100,    //token刷新有效时间
		IdentityKey:           "username",              // 用户关键字
		TokenLookup:           "header: Authorization", // 捕抓请求的指定数据
		TokenHeadName:         "jwt",                   // token 头名称
		TimeFunc:              time.Now,
		Authenticator:         SimpleAuthenticator,   //登录验证
		LoginResponse:         LoginResponse,         //登录返回token
		RefreshResponse:       RefreshResponse,       //刷新token
		Unauthorized:          Unauthorized,          //未登录返回
		IdentityHandler:       IdentityHandler,       //返回数据给Authorizator
		PayloadFunc:           PayloadFunc,           //将Authenticator返回的内容记录到jwt
		Authorizator:          Authorizator,          //接收IdentityHandler数据并判断权限
		HTTPStatusMessageFunc: HTTPStatusMessageFunc, //错误处理
	})
	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}
	GfJWTMiddleware = authMiddleware
}

// GetLoginCryptoKey 获取登录的加密key
//
// @Summary 获取登录的加密key
// @Description 获取登录的加密key
// @Tags auth
// @Success 200 {string} string	"ok"
// @router /loginkey [get]
func GetLoginCryptoKey(r *ghttp.Request) {
	kid := r.Session.Id()
	ck := common.GenCryptoKey(kid)
	//glog.Debug("kid:" + kid)
	Success(r, ck)
}

//
// @Summary 登出
// @Description 登出
// @Tags auth
// @Success 200 {string} string	"ok"
// @router /logout [post]
func Logout(r *ghttp.Request) {
	Success(r, "success")
}

//
// @Summary 登陆
// @Description 登陆
// @Tags auth
// @Param   Login  body api_model.Login true "Login"
// @Success 200 {string} string	"ok"
// @router /login [post]
func SwaggerLogin(r *ghttp.Request) {
	Success(r, "success")
}

//
// @Summary 刷新token
// @Description 刷新token
// @Tags auth
// @Success 200 {string} string	"ok"
// @router /refresh_token [get]
func SwaggerRefreshToken(r *ghttp.Request) {
	Success(r, "success")
}
