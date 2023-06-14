package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"server/config"
)

func handlerLogrus() {
	fmt.Println("logrus handler......")
}

//log file
var fileLog *os.File

//Logrus日志
func init() {
	//设置日志级别
	logrus.SetLevel(logrus.DebugLevel)
	//记录调用者
	logrus.SetReportCaller(false)
	//设置日志格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		DisableHTMLEscape: true,
	})

	//钩子
	logrus.RegisterExitHandler(handlerLogrus)

	//日志输出
	logFilePath := "./logrus.log"
	if config.Config.Log.LogFilePath != "" {
		logFilePath = config.Config.Log.LogFilePath
	}
	if fileLog, _ = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); fileLog != nil {
		logrus.SetOutput(fileLog)
	} else {
		logrus.SetOutput(os.Stdout)
	}
}

//Error
func Error(ctx context.Context, args string) {
	logrus.WithFields(logrus.Fields{
		"TraceId": ctx.Value("trace_id"),
		"UserId":  ctx.Value("user_id"),
		"IP":      ctx.Value("client_ip"),
	}).Error(args)
}

//func Errorf(ctx context.Context, format string, args ...any) {
//	logrus.WithFields(logrus.Fields{
//	"TraceId": ctx.Value("trace_id"),
//	"IP":      ctx.Value("client_ip"),
//	//})
//	logrus.Errorf(format, args...)
//}

//Warning
func Warning(ctx context.Context, args string) {
	logrus.WithFields(logrus.Fields{
		"TraceId": ctx.Value("trace_id"),
		"UserId":  ctx.Value("user_id"),
		"IP":      ctx.Value("client_ip"),
	}).Error(args)
}

//func Warningf(format string, args ...any) {
//	logrus.Warningf(format, args...)
//}

//Debug
func Debug(ctx context.Context, args string) {
	logrus.WithFields(logrus.Fields{
		"TraceId": ctx.Value("trace_id"),
		"UserId":  ctx.Value("user_id"),
		"IP":      ctx.Value("client_ip"),
	}).Debug(args)
}

//func Debugf(format string, args ...any) {
//	logrus.Debugf(format, args...)
//}

//Info
func Info(ctx context.Context, args string) {
	logrus.WithFields(logrus.Fields{
		"TraceId": ctx.Value("trace_id"),
		"UserId":  ctx.Value("user_id"),
		"IP":      ctx.Value("client_ip"),
	}).Debug(args)
}

//func Infof(format string, args ...any) {
//	logrus.Infof(format, args...)
//}

//关闭文件句柄
func FinishClear() {
	fileLog.Close()
}
