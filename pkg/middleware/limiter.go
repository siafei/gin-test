package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/siafei/gin-test/pkg/limiter"
	"github/siafei/gin-test/pkg/response"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		fmt.Println(key)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				res := response.NewResponse(c)
				res.ToErrorResponse(1,"请求过多")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
