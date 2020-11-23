package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.Use()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080
}
