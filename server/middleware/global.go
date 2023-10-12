package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/core/logger"
	"server/library/command"
	"time"
)

// Global 中间件: 全局中间件
func Global() gin.HandlerFunc {
	return func(c *gin.Context) {
		//start time
		start := time.Now()

		//gin context
		userId := c.Request.Header.Get("user_id")   //user id
		traceId := c.Request.Header.Get("trace_id") //链路追踪ID
		if traceId == "" {
			traceId = command.GenTraceID() //生成链路追踪ID
		}
		c.Set("client_ip", c.ClientIP()) //client ip
		c.Set("user_id", userId)         //user id
		c.Set("trace_id", traceId)       //链路追踪ID

		//before request
		c.Next()
		//after request

		//执行耗时(ms)
		etime := time.Since(start).Milliseconds()
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUri := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()

		//日志格式
		logger.Debug(c, fmt.Sprintf(`[URI:%s | Status Code:%d | Execution Time(ms):%d | Method:%s]`, reqUri, statusCode, etime, reqMethod))
	}
}
