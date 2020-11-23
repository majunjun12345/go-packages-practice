package main

import (
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	//rate-limit 限流中间件
	lmt := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Second * 1,
	})

	lmt.SetMessage("服务繁忙，请稍后再试...")
	// lmt.SetOnLimitReached(func(w http.ResponseWriter, r *http.Request) { fmt.Println("A request was rejected") })
	// lmt.ExecOnLimitReached(w http.ResponseWriter, r *http.Request)

	// lmt.SetTokenBucketExpirationTTL(time.Second * 5)
	// lmt.SetBasicAuthExpirationTTL(time.Second * 5)
	// lmt.SetHeaderEntryExpirationTTL(time.Second * 5)

	tollbooth.LimitByKeys(lmt, []string{"127.0.0.1", "/"})

	r.GET("/ping", tollbooth_gin.LimitHandler(lmt), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ping": "pong",
		})
	})
	r.Run("127.0.0.1:8082")
}
