package main

import (
	"github.com/gogf/gf/g"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/logger"
	"github.com/hailaz/gadmin/library/timer"
	"github.com/hailaz/gadmin/router"
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

// @host
// @BasePath /
func main() {
	s := g.Server()
	s.SetIndexFolder(false)
	s.SetIndexFiles([]string{"index.html"})
	s.SetServerRoot(".")
	// 初始化路由
	router.InitRouter(s)

	s.SetPort(g.Config().GetInt("port", 8080))
	s.Run()

}
