package logger

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/glog"
)

// InitLogger 初始化日志设置
//
// createTime:2019年05月13日 09:46:02
// author:hailaz
func InitLogger() {
	path := g.Config().GetString("logpath", "log")
	glog.SetPath(path)
	glog.SetLevel(glog.LEVEL_ALL)
	glog.SetFlags(glog.F_TIME_STD | glog.F_FILE_SHORT)
}
