package api

import (
	"github.com/gogf/gf-jwt"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/utils"
	"time"
)

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
	//glog.Debugfln("user:%v ,method:%v ,path:%v", data, method, path)
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
		Token  string `json:"token"`
		Expire string `json:"expire"`
	}
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
			return nil, err
		}
		if u.Password == utils.EncryptPassword(password) {
			return g.Map{
				"username": u.UserName,
				"id":       u.Id,
			}, nil
		}

	}

	return nil, jwt.ErrFailedAuthentication
}
