package timer

import (
	"github.com/gogf/gf/os/gtimer"
	"github.com/hailaz/gadmin/app/model"
	"time"
)

// InitTimer 初始化定时任务
//
// createTime:2019年04月24日 14:50:34
// author:hailaz
func InitTimer() {
	gtimer.Add(time.Minute, model.InitCasbin)
}
