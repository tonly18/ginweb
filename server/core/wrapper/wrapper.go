package wrapper

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"server/core/iface"
	"server/core/request"
)

func HandlerFuncWrapper(handler iface.IWrapperHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		//gin context
		userId := cast.ToInt(c.GetString("user_id")) //user id
		clientIp := c.ClientIP()                     //client ip
		traceId := c.GetString("trace_id")           //链路追踪ID

		//handle request
		req := request.NewRequest(c, userId, clientIp, traceId)
		handler.PreHandler(req)
		resp, err := handler.Handler(req)
		handler.PostHandler(req)

		//response data
		if err != nil {
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
