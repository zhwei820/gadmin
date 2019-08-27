package context_log

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

// 是否开启debug模式
var defaultEnableDebugmode = false

// SysLog 调试模式打印系统日志
func SysLog(v ...interface{}) {
	if defaultEnableDebugmode {
		log.Output(2, "[SYS] "+fmt.Sprintln(v...))
	}
}

// SysLogf 调试模式打印格式化系统日志
func SysLogf(format string, v ...interface{}) {
	if defaultEnableDebugmode {
		data := fmt.Sprintf(format, v...)
		log.Output(2, "[SYS] "+data)
	}
}

// 日志等级
//###背景
//现有的log库都是以单行日志为单位的，即一条日志的落地与否只与该条日志的级别和程序配置的日志级别
//有关(如：程序配置的日志级别为INFO，则级别高于等于INFO的日志都会落地)。实践中一般测试环境配置为DEBUG级别，
//现网环境配置为WARNING或INFO级别。但是后台逻辑层服务响应每个请求都需要并发或串行请求多个后端服务，如果按
//照常规做法只打印出错的那一条日志对定位问题很不利，原因：1.由于是以单条log为单位的，所以同一请求的log会分
//散在日志文件的不同位置，不连续，查看log不方便；2.导致上下文缺失，经常需要让用户重试，帮忙复现问题，但并不
//是每个用户都有时间陪你玩。当然如果不拍把磁盘打爆、被老板叼线上也可以开启DEBUG。。。
//###目标
//既然传统的以单条日志为单位的打log方式不方便线上定位问题，那是否可以以单次请求为单位，只要这次请求所产生的
//log中任何一条级别大于程序配置的日志级别，就将所有日志都落地(包括DEBUG、INFO等)；本模块正式基于这样的目的
//设计的，将每个请求对应的所有log作为一个group，每个请求对应一个logger对象，在处理请求的过程中产生的log不
//会立刻落地而是缓存在logger对象中，等请求处理结束再根据group的日志级别(group的级别由其缓存的所以日志中级
//别最高的一条决定)决定是否需要落地，并且落地操作为一次批量写入；这样就解决了同一请求的日志不连续和出问题时
//上线文缺失的问题。
const (
	LogLevelNull    = 0
	LogLevelTrace   = 1
	LogLevelDebug   = 2
	LogLevelInfo    = 3
	LogLevelWarning = 4
	LogLevelError   = 5
	LogLevelFatal   = 6
)

// 客户端来源
const (
	SourceIOS     = 1
	SourceAndroid = 2
	SourceWeb     = 3
	SourceUnknown = 99
)

// MultiFuncContext 多功能的context接口, 提供超时控制, 提供日志, 是所有going里的context都实现的接口
type MultiFuncContext interface {
	context.Context
	Trace(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Error(format string, v ...interface{})
}

// Context 携带log，uin，cmd，subcmd的请求上下文
type Context struct {
	context.Context
	*log.Logger
	Seq           uint32
	LogLevel      int         //日志打印等级，每次new一个context时，都必须设置这个
	writer        io.Writer
	buffer        *bytes.Buffer
	startTime     time.Time
	level         int
	NoResponse    bool
}

// NewContext 新建一个包含日志的ctx，
func NewContext(ctx context.Context, w io.Writer) *Context {
	newCtx := Context{
		Context:   ctx,
		writer:    w,
		startTime: time.Now(),
	}
	newCtx.buffer = new(bytes.Buffer)
	newCtx.Logger = log.New(newCtx.buffer, "", log.LstdFlags|log.Lshortfile)
	return &newCtx
}

// Trace 打印trace日志 log trace, no uls
func (ctx *Context) Trace(format string, v ...interface{}) {
	if ctx.level < LogLevelTrace {
		ctx.level = LogLevelTrace
	}
	var buffer bytes.Buffer
	buffer.WriteString("[TRACE] ")
	buffer.WriteString(fmt.Sprintf(format, v...))
	ctx.Output(2, buffer.String())
}

// Debug 打印debug日志 log debug, uls debug
func (ctx *Context) Debug(format string, v ...interface{}) {
	data := fmt.Sprintf(format, v...)
	//(&uls.Logger{
	//	Seq:           ctx.Seq,
	//	ClientIP:      ctx.ClientAddr,
	//	ClientVersion: ctx.ClientVersion,
	//	ErrCode:       ctx.ErrCode,
	//	Level:         uls.LogLevelDebug,
	//	Data:          data,
	//}).Report()

	if ctx.level < LogLevelDebug {
		ctx.level = LogLevelDebug
	}
	var buffer bytes.Buffer
	buffer.WriteString("[DEBUG] ")
	buffer.WriteString(data)
	ctx.Output(2, buffer.String())
}

// Info 打印info日志 log info, uls info
func (ctx *Context) Info(format string, v ...interface{}) {
	data := fmt.Sprintf(format, v...)

	if ctx.level < LogLevelInfo {
		ctx.level = LogLevelInfo
	}
	var buffer bytes.Buffer
	buffer.WriteString("[INFO] ")
	buffer.WriteString(data)
	ctx.Output(2, buffer.String())
}

// Warning 打印warning日志 log warning, uls warning
func (ctx *Context) Warning(format string, v ...interface{}) {
	data := fmt.Sprintf(format, v...)

	if ctx.level < LogLevelWarning {
		ctx.level = LogLevelWarning
	}
	var buffer bytes.Buffer
	buffer.WriteString("[WARN] ")
	buffer.WriteString(data)
	ctx.Output(2, buffer.String())
}

// Error 打印error日志 log error, uls error
func (ctx *Context) Error(format string, v ...interface{}) {
	data := fmt.Sprintf(format, v...)

	if ctx.level < LogLevelError {
		ctx.level = LogLevelError
	}
	var buffer bytes.Buffer
	buffer.WriteString("[ERROR] ")
	buffer.WriteString(data)
	ctx.Output(2, buffer.String())
}

// Fatal 打印fatal日志 log fatal, uls error
func (ctx *Context) Fatal(format string, v ...interface{}) {
	data := fmt.Sprintf(format, v...)

	if ctx.level < LogLevelFatal {
		ctx.level = LogLevelFatal
	}
	var buffer bytes.Buffer
	buffer.WriteString("[FATAL] ")
	buffer.WriteString(data)
	ctx.Output(2, buffer.String())
}

// Level 返回当前日志最高级别 return log level
func (ctx *Context) Level() int {
	return ctx.level
}

// WriteLog 请求结束后刷新日志到文件 write all log into file when request finish
func (ctx *Context) WriteLog() {
	if ctx.level >= ctx.LogLevel && ctx.buffer.Len() > 0 {
		ctx.writer.Write(ctx.buffer.Bytes())
		ctx.writer.Write([]byte(fmt.Sprintf("==> %v\n", ctx.Cost().Seconds())))
	}
	ctx.buffer.Reset()
	ctx.level = LogLevelNull
}

// Now 返回请求进入时间 return enter time, no need to call time.Now() every time , every where
func (ctx *Context) Now() time.Time {
	return ctx.startTime
}

// Cost 返回当前耗时 return cost time
func (ctx *Context) Cost() time.Duration {
	return time.Now().Sub(ctx.startTime)
}

// String 打印ctx信息 防止高并发下多个goroutine打印整个ctx导致crash问题
func (ctx *Context) String() string {
	return fmt.Sprintf("seq[%d]: ", ctx.Seq)
}


// IsDone 判断ctx是否已经失效，包括超时，被取消
func (ctx *Context) IsDone() bool {
	select {
	case <-ctx.Done():
		ctx.Debug("context is done, err:%s", ctx.Err().Error())
		return true
	default:
	}

	return false
}
