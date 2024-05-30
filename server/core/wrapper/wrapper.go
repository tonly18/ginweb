package wrapper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"server/core/iface"
	"server/core/logger"
	"server/core/request"
)

func HandlerFuncWrapper(handler iface.IWrapperHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(c, fmt.Sprintf(`[wrapper panic] Error(1): %v`, err))
				for i := 1; i < 20; i++ {
					if pc, file, line, ok := runtime.Caller(i); ok {
						logger.Error(c, fmt.Sprintf(`[wrapper panic] goroutine:%v, file:%s, function name:%s, line:%d`, pc, file, runtime.FuncForPC(pc).Name(), line))
					}
				}
				logger.Error(c, fmt.Sprintf(`[wrapper panic] Error(2): %v`, err))
			}
		}()

		//handle request
		req := request.NewRequest(c)
		handler.PreHandler(req)
		resp, err := handler.Handler(req)
		handler.PostHandler(req)

		if err != nil {
			for _, e := range err.GetStack() {
				logger.Error(req, fmt.Sprintf(`[%d] message:%v, error:%v`, e.GetCode(), e.GetMsg(), e.GetRawError()))
			}
			c.JSON(http.StatusOK, gin.H{
				"code": err.GetCode(),
				"data": nil,
				"msg":  err.GetMsg(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"data": resp.Data,
				"msg":  "",
			})
		}
	}
}
