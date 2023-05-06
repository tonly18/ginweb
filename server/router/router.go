package router

import (
	"github.com/gin-gonic/gin"
	"server/library/middleware"
)

//RouterInit 初始化路由
func RouterInit() *gin.Engine {
	r := gin.New()

	//添加中间件:全局中间件
	r.Use(middleware.Global())
	r.Use(middleware.RecoverPanic())

	//路由组:须验证状态
	routerGroupInit(r)

	//return
	return r
}
