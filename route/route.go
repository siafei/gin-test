package route

import (
	"github.com/gin-gonic/gin"
	"github/siafei/gin-test/app/controller"
	"github/siafei/gin-test/pkg/limiter"
	"github/siafei/gin-test/pkg/middleware"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/api/user",	//限流的路由
	FillInterval: time.Second,	//间隔多久时间放 N 个令牌
	Capacity:     10,	//令牌桶的容量
	Quantum:      10,	//每次到达间隔时间后所放的具体令牌数量
})

func NewRoute() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.Translations())
	r.Use(middleware.RateLimiter(methodLimiters))
	api := r.Group("api/")
	{
		user_route := api.Group("user")
		{
			user := controller.UserController{}
			user_route.GET("", user.GetUsers)
		}

	}
	return r
}
