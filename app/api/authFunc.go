package api

import (
	"encoding/base64"
	"github.com/gogf/gf-jwt"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/app/service"
	"github.com/hailaz/gadmin/library/common"
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

// Authenticator 登录验证
//
// createTime:2019年05月13日 10:00:22
// author:hailaz
func Authenticator(r *ghttp.Request) (interface{}, error) {
	data := r.GetJson()
	name := data.GetString("username")
	pwd := data.GetString("password")
	kid := data.GetString("kid")

	if ck := common.GetCryptoKey(kid); ck != nil {
		if gtime.Second()-ck.TimeStamp >= 5 { //加密key超时时间
			return nil, jwt.ErrFailedAuthentication
		}
		//glog.Debugfln("%v", ck.Id)
		//glog.Debugfln("%v", ck.Key)
		//glog.Debugfln("%v %v", name, pwd)
		decodePwd, err := base64.StdEncoding.DecodeString(pwd)
		if err != nil {
			return nil, err
		}
		decryptPwd, _ := common.RsaDecrypt(decodePwd, []byte(ck.PrivateKey))
		//glog.Debug(string(decryptPwd))
		password := string(decryptPwd)
		//glog.Debugfln("%v %v", name, password)
		if password != "" {
			u, err := model.GetUserByName(name)
			if err != nil {
				return nil, err
			}
			if u.Password == service.EncryptPassword(password) {
				return g.Map{
					"username": u.UserName,
					"id":       u.Id,
				}, nil
			}

		}
	}

	return nil, jwt.ErrFailedAuthentication
}
