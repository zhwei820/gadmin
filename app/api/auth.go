package api

import (
	"encoding/base64"
	"github.com/hailaz/gadmin/app/service"
	"time"

	"github.com/gogf/gf-jwt"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/hailaz/gadmin/app/model"
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
		Timeout:               time.Minute * 10,        //token有效时间
		MaxRefresh:            time.Minute * 10,        //token刷新有效时间
		IdentityKey:           "username",              // 用户关键字
		TokenLookup:           "header: Authorization", // 捕抓请求的指定数据
		TokenHeadName:         "gadmin",                // token 头名称
		TimeFunc:              time.Now,
		Authenticator:         Authenticator,         //登录验证
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
// createTime:2019年04月24日 13:57:34
// author:hailaz
func GetLoginCryptoKey(r *ghttp.Request) {
	kid := r.Session.Id()
	ck := common.GenCryptoKey(kid)
	//glog.Debug("kid:" + kid)
	Success(r, ck)
}

func Logout(r *ghttp.Request) {
	Success(r, "success")
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
