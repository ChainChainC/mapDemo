package main

import (
	"mapDemo/common"

	"github.com/gin-gonic/gin"
)

func init() {
	common.NewRedisClientApp()
}

func main() {

	r := gin.Default()
	r = CollectRoute(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
