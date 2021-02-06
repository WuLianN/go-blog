package routers

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/WuLianN/go-blog/internal/routers/api"
	"github.com/WuLianN/go-blog/internal/middleware"
	"github.com/WuLianN/go-blog/global"
	"github.com/WuLianN/go-blog/pkg/limiter"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth", // 自定义键值对名称
		FillInterval: time.Second, // 间隔多久时间放 Quantum 个令牌
		Capacity:     10, // 令牌桶的容量
		Quantum:      10, // 每次到达间隔时间后所放的具体令牌数量
	},
)


func SetupRouter() *gin.Engine {
    r := gin.Default()

	// 接口限流控制
	r.Use(middleware.RateLimiter(methodLimiters))
	// 统一超时管理
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))

	r.GET("/auth", api.GetAuth)
    
	apiv1 := r.Group("api/v1")
	// apiv1.Use(middleware.JWT())
	{
		apiv1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	
    
    return r
}