package tests

//
//import (
//	"github.com/gogf/gf/frame/g"
//	"github.com/zhwei820/gadmin/router"
//	"github.com/zhwei820/gogf-swagger"
//
//	//"github.com/zhwei820/gogf-swagger"
//	"github.com/zhwei820/gogf-swagger/swaggerFiles"
//)
//
//func init()  {
//	s := g.Server()
//	s.SetLogStdout(true)
//	s.SetAccessLogEnabled(true)
//	// 初始化路由
//	router.InitRouter(s)
//
//	url := gogfSwagger.URL("http://localhost:8199/swagger/doc.json") //The url pointing to API definition
//	s.BindHandler("/swagger/*any", gogfSwagger.WrapHandler(swaggerFiles.Handler, url))
//
//	s.SetPort(g.Config().GetInt("port", 8199))
//	s.Run()
//
//}
