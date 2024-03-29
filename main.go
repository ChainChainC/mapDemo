package main

import (
	"mapDemo/common"
	"mapDemo/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	common.NewRedisClientApp()
	log.SetLevel(logrus.InfoLevel)
	log.SetReportCaller(true)
	// 默认情况下，日志输出到io.Stderr
}

func main() {

	r := gin.Default()
	r = CollectRoute(r)
	r.Use(middleware.LoggerMiddleWare())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":30022") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
