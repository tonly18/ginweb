package router

import (
	"github.com/gin-gonic/gin"
	"server/library/middleware"
)

//routerGroupInit 初始化路由组
func routerGroupInit(r *gin.Engine) {
	//路由组: 须验证状态
	rGroupVerify := r.Group("/v", middleware.LoginVerify())
	{
		newRouterGroupPath(rGroupVerify).verifyLogin()
	}

	//路由组: 不须验证状态
	rGroupVerifyNot := r.Group("/n", middleware.LoginNotVerify())
	{
		newRouterGroupPath(rGroupVerifyNot).verifyLoginNot()
	}
}
