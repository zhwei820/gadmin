package main

import (
	"github.com/gogf/gf/g"
	"github.com/hailaz/gadmin/app/model"
	_ "github.com/hailaz/gadmin/docs"
	"github.com/hailaz/gadmin/library/logger"
	"github.com/hailaz/gadmin/library/timer"
	"github.com/hailaz/gadmin/router"
	"github.com/zhwei820/gogf-swagger"
	"github.com/zhwei820/gogf-swagger/swaggerFiles"
)

func init() {
	// 设置默认配置文件，默认的 config.toml 将会被覆盖
	g.Config().SetFileName("config.toml")
	// 初始化数据库
	model.InitModel()
	// 初始化日志
	logger.InitLogger()

	timer.InitTimer()

}

// @title Swagger Example API
// @version 1.0
// @description This is a hello server .
// @termsOfService http://hello.io/terms/

// @contact.name hello
// @contact.url http://www.hello.io
// @contact.email hello@hello.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host
// @BasePath /
func main() {
	s := g.Server()
	s.SetLogStdout(true)
	s.SetAccessLogEnabled(true)
	// 初始化路由
	router.InitRouter(s)

	url := gogfSwagger.URL("http://localhost:8199/swagger/doc.json") //The url pointing to API definition
	s.BindHandler("/swagger/*any", gogfSwagger.WrapHandler(swaggerFiles.Handler, url))

	s.SetPort(g.Config().GetInt("port", 8199))
	s.Run()

}
