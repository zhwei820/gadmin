package context_log

import (
	"context"
	"github.com/gogf/gf/os/glog"
	"testing"
)

func work() {

	w, _ := NewFileWriter("log.log", 1024*1024*50, 30)
	ctx := NewContext(context.Background(), w, "test", glog.LEVEL_ALL)
	ctx.Debug("hi i am debug")
	ctx.Info("hi i am Info")
	ctx.Error("hi i am Error")
	ctx.Error("hi i am Error")
	ctx.WriteLog()
}

func TestContext_WriteLog(t *testing.T) {
	for ii := 0; ii < 10; ii += 1 {
		go work()
	}
	select {}
}
