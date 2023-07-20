package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/core/logger"
	"server/library/command"
	"time"
)

//Global 中间件: 全局中间件
func Global() gin.HandlerFunc {
	return func(c *gin.Context) {
		//gin context
		traceId := c.Request.Header.Get("trace_id") //链路追踪ID
		userId := c.Request.Header.Get("user_id")   //user id
		if traceId == "" {
			traceId = command.GenTraceID() //生成链路追踪ID
		}
		c.Set("trace_id", traceId)       //链路追踪ID
		c.Set("user_id", userId)         //user id
		c.Set("client_ip", c.ClientIP()) //client ip

		//开始时间
		startTime := time.Now().UnixMilli()

		//before request
		c.Next()
		//after request

		//结束时间
		endTime := time.Now().UnixMilli()
		//执行时间
		latencyTime := endTime - startTime
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUri := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//后续根据需求再做修改...

		//日志格式
		logger.Debug(c, fmt.Sprintf(`[URI:%s | Status Code:%d | Execution Time(ms):%d | Method:%s]`, reqUri, statusCode, latencyTime, reqMethod))
	}
}
