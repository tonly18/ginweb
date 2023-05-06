package middleware

import (
	"github.com/gin-gonic/gin"
)

//LoginVerify 中间件: 验证登录状态
func LoginVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		//验证登录
		//uid, sid, err := checkLogin(c)

		// before request
		c.Next()
		// after request
	}
}

//LoginNotVerify 中间件: 不验证登录状态
func LoginNotVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: fix me
	}
}

//checkLogin 中间件: 验证登录
func checkLogin(c *gin.Context) (string, string, error) {
	return "", "", nil
}
