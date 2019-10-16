package common

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/zhwei820/gadmin/app/model"
	"strings"
)

var RouterMap = make(map[string]model.RolePolicy)

// BindGroup 绑定分组路由
//
// createTime:2019年04月29日 16:45:55
// author:hailaz
func BindGroup(s *ghttp.Server, path string, items []ghttp.GroupItem) {
	g := s.Group(path)
	g.Bind(items)
	for _, item := range items {
		glog.Debug(gconv.String(item[1]))
		if gconv.String(item[0]) == "REST" { //rest api
			addPolicy("system", path+gconv.String(item[1]), model.ACTION_GET)
			addPolicy("system", path+gconv.String(item[1]), model.ACTION_POST)
			addPolicy("system", path+gconv.String(item[1]), model.ACTION_PUT)
			addPolicy("system", path+gconv.String(item[1]), model.ACTION_DELETE)
		} else {
			addPolicy("system", path+gconv.String(item[1])+"*", GetAction(gconv.String(item[0])))
		}
	}

}

// addPolicy 记录需要系统路由
//
// createTime:2019年04月29日 17:18:25
// author:hailaz
func addPolicy(role, path, act string) {
	RouterMap[fmt.Sprintf("%v:%v", path, act)] = model.RolePolicy{Role: role, Path: path, Act: act}
}

func GetAction(act string) string {
	acts := strings.Split(strings.Split(strings.Split(act, ";")[0], ":")[0], ",")
	action := ""
	for _, v := range acts {
		if v == "All" || v == "REST" {
			return model.ACTION_ALL
		}
		if action == "" {
			action += strings.ToUpper("(" + v + ")")
		} else {
			action += strings.ToUpper("|(" + v + ")")
		}

	}
	return action
}
