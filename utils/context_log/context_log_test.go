package context_log

import (
	"context"
	"testing"
)

func work() {

	w, _ := NewFileWriter("log.log", 9999, 3)
	ctx := NewContext(context.Background(), w, "test", 2)
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
