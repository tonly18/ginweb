package wrapper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"server/core/iface"
	"server/core/logger"
	"server/core/request"
	"strconv"
)

func HandlerFuncWrapper(handler iface.IWrapperHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(c, fmt.Sprintf(`[wrapper panic] Error(1): %v`, err))
				logger.Error(c, fmt.Sprintf(`[wrapper panic] User Id: %v, Client IP: %v`, c.GetString("user_id"), c.GetString("client_ip")))
				for i := 1; i < 20; i++ {
					if pc, file, line, ok := runtime.Caller(i); ok {
						logger.Error(c, fmt.Sprintf(`[wrapper panic] goroutine:%v, file:%s, function name:%s, line:%d`, pc, file, runtime.FuncForPC(pc).Name(), line))
					}
				}
				logger.Error(c, fmt.Sprintf(`[wrapper panic] Error(2): %v`, err))
			}
		}()

		//gin context
		clientIp := c.GetString("client_ip")              //client ip
		userId, _ := strconv.Atoi(c.GetString("user_id")) //user id
		traceId := c.GetString("trace_id")                //链路追踪ID

		//handle request
		req := request.NewRequest(c, userId, clientIp, traceId)
		handler.PreHandler(req)
		resp, err := handler.Handler(req)
		handler.PostHandler(req)

		//response data
		if err != nil {
			//error log
			for _, e := range err.GetErrorStack() {
				logger.Error(req, fmt.Sprintf(`[%d] message:%v, error:%v`, e.GetCode(), e.GetMsg(), e.GetErr()))
			}
			//gin
			c.JSON(http.StatusOK, gin.H{
				"code": err.GetCode(),
				"msg":  err.GetMsg(),
				"data": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": resp.Code,
				"msg":  resp.Msg,
				"data": resp.Data,
			})
		}
	}
}
