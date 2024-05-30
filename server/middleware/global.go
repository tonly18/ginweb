package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/core"
	"server/core/logger"
	"server/library/command"
	"time"
)

// Global 中间件: 全局中间件
func Global() gin.HandlerFunc {
	return func(c *gin.Context) {
		//start time
		startTime := time.Now()

		//gin context
		traceId := c.Request.Header.Get(core.TraceID) //链路追踪ID
		if traceId == "" {
			traceId = command.GenTraceID() //生成链路追踪ID
		}
		c.Set(core.ClientIP, c.ClientIP()) //client ip
		c.Set(core.TraceID, traceId)       //链路追踪ID

		//before request
		c.Next()
		//after request

		//日志格式
		logger.Infof(c, fmt.Sprintf(`[URI:%s | Method:%s | Status Code:%d | Execution Time(ms):%d]`, c.Request.RequestURI, c.Request.Method, c.Writer.Status(), time.Since(startTime).Milliseconds()))
	}
}
