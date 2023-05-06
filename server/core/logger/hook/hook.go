package hook

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TraceIdHook struct {
	GinCtx *gin.Context
}

func NewTraceIdHook(c *gin.Context) logrus.Hook {
	hook := TraceIdHook{
		GinCtx: c,
	}
	return &hook
}

func (hook *TraceIdHook) Fire(entry *logrus.Entry) error {
	// 设置链路Id字段至日志中
	entry.Data["UUID"] = hook.GinCtx.GetString("trace_id")
	return nil
}

func (hook *TraceIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
