package wrapper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"server/core/iface"
	"server/core/logger"
	"server/core/request"
)

func HandlerFuncWrapper(handler iface.IWrapperHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		//gin context
		userId := cast.ToInt(c.GetString("user_id")) //user id
		clientIp := c.GetString("client_ip")         //client ip
		traceId := c.GetString("trace_id")           //链路追踪ID

		//handle request
		req := request.NewRequest(c, userId, clientIp, traceId)
		handler.PreHandler(req)
		resp := handler.Handler(req)
		handler.PostHandler(req)
		if resp.Type == 1 {
			logger.Error(req, fmt.Sprintf(`[user id:%v, code:%v, data:%v, message:%v]`, userId, resp.Code, resp.Data, resp.Msg))
		}

		//result
		c.JSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
			"data": resp.Data,
		})
	}
}
