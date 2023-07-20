package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//LoginVerify 中间件: 验证登录状态
func LoginVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		//验证登录
		_, _, err := checkLogin(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 1000,
				"msg":  nil,
				"data": nil,
			})
		}

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
func checkLogin(c *gin.Context) (int, int, error) {
	//TODO: fix me

	return 0, 0, nil
}
