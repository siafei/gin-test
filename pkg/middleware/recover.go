package middleware

import (
	"github.com/gin-gonic/gin"
	"github/siafei/gin-test/global"
	"github/siafei/gin-test/pkg/response"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().ErrorOf("panic recover err: %v", err)
				response.NewResponse(c).ToErrorResponse(1,err.(string))
				c.Abort()
			}
		}()
		c.Next()
	}
}
