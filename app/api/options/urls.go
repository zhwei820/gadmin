package options

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/zhwei820/gadmin/common"
)

// Init 初始化rbac
//
// createTime:2019年04月25日 09:24:06
// author:hailaz
func InitRouter(s *ghttp.Server) {
	optionsCtrl := new(OptionsController)

	// user
	common.BindGroup(s, "/options", []ghttp.GroupItem{
		{"GET", "/", optionsCtrl},
	})
}
