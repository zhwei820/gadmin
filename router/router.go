package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/hailaz/gadmin/app/api/base"
	"github.com/hailaz/gadmin/app/api/options"
	"github.com/hailaz/gadmin/app/api/rbac"
	"github.com/hailaz/gadmin/app/service"
	"github.com/hailaz/gadmin/common"
	"github.com/hailaz/gadmin/utils"
	"github.com/hailaz/gadmin/utils/context_log"
	"github.com/satori/go.uuid"
	"io"
	"strings"

	"github.com/gogf/gf/net/ghttp"
)

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

	rbac.InitRouter(s)    // rbac router
	options.InitRouter(s) // options router

	service.ReSetPolicy("system", common.RouterMap)
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

	r.Response.CORS(ghttp.CORSOptions{
		AllowOrigin:      "*",
		AllowMethods:     ghttp.HTTP_METHODS,
		AllowCredentials: "true",
		MaxAge:           3628800,
		AllowHeaders:     "*",
	}) //开启跨域

	if r.Request.Method == "OPTIONS" {
		return
	}
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
	base.GfJWTMiddleware.MiddlewareFunc()(r) //鉴权中间件
	// or error handling
}
