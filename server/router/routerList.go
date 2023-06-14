package router

import (
	"github.com/gin-gonic/gin"
	"server/controller"
	"server/core/wrapper"
)

//routerGroupPath 路由设置
type routerGroupPath struct {
	ginRouterGroup *gin.RouterGroup
}

func newRouterGroupPath(routerGroup *gin.RouterGroup) *routerGroupPath {
	return &routerGroupPath{
		ginRouterGroup: routerGroup,
	}
}

//verifyLoginNot 不需验证登录状态
func (r routerGroupPath) verifyLoginNot() {
	r.ginRouterGroup.GET("/v1/test", wrapper.HandlerFuncWrapper(&controller.TestHandler{}))
	r.ginRouterGroup.GET("/v1/stu", wrapper.HandlerFuncWrapper(&controller.StudentHandler{}))
}

//verifyLogin 须验证登录状态
func (r *routerGroupPath) verifyLogin() {
	//r.ginRouterGroup.GET("/v1/test", wrapper.HandlerFuncWrapper(&controller.TestHandler{}))
}
