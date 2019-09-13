package base

import (
	"errors"
	"fmt"
	"github.com/gogf/gf-jwt"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/utils/crypt"
	"time"
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

func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

func Authorizator(data interface{}, r *ghttp.Request) bool {
	method := r.Method
	path := r.URL.Path
	glog.Debugf("user:%v ,method:%v ,path:%v\n", data, method, path)
	return model.Enforcer.Enforce(data, path, method)
}

// IdentityHandler sets the identity for JWT.
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return claims["username"]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(r *ghttp.Request, code int, message string) {
	Fail(r, code, message)
}

func HTTPStatusMessageFunc(e error, r *ghttp.Request) string {
	glog.Debug(e.Error())
	switch e.Error() {
	case "Token is expired":
		return "token超时"
	}
	return e.Error()
}

// LoginResponse is used to define customized login-successful callback function.
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	var tk struct {
		Token  string   `json:"token"`
		Expire string   `json:"expire"`
		Perms  []string `json:"perms"` // policys
	}
	policys := model.Enforcer.GetPermissionsForUser(r.GetParam("username").String())
	Perms := make([]string, 0)
	for _, item := range policys {
		perm := fmt.Sprintf("%v:%v", item[1], item[2])
		if perm == "*:(GET)|(POST)|(PUT)|(DELETE)|(PATCH)|(OPTIONS)|(HEAD)" {
			perm = "superuser"
		}
		Perms = append(Perms, perm)
	}

	tk.Perms = Perms
	tk.Token = GfJWTMiddleware.TokenHeadName + " " + token
	tk.Expire = expire.Format(time.RFC3339)
	Success(r, tk)
}

// RefreshResponse is used to get a new token no matter current token is expired or not.
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	var tk struct {
		Token  string `json:"token"`
		Expire string `json:"expire"`
	}
	tk.Token = GfJWTMiddleware.TokenHeadName + " " + token
	tk.Expire = expire.Format(time.RFC3339)
	Success(r, tk)
}

// 简单 Authenticator 登录验证
func SimpleAuthenticator(r *ghttp.Request) (interface{}, error) {
	data := r.GetJson()
	name := data.GetString("username")
	password := data.GetString("password")

	//glog.Debugfln("%v %v", name, password)
	if password != "" {
		u, err := model.GetUserByName(name)
		if err != nil {
			return nil, errors.New("用户名, 密码错误")
		}
		if u.Password == crypt.EncryptPassword(password) {
			r.SetParam("username", u.Username)
			return g.Map{
				"username": u.Username,
				"id":       u.Id,
			}, nil
		}

	}

	return nil, jwt.ErrFailedAuthentication
}
