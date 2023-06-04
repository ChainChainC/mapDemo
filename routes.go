package main

import (
	"mapDemo/controller"
	. "mapDemo/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	Logger().Debug("start initializing routes...")

	// 解决跨域问题
	r.Use(CORSMiddleware())

	// 房间
	room := r.Group("api/room")
	{
		room.POST("/NewRoom", controller.NewRoom)
	}
	player := r.Group("api/player")
	{
		player.POST("/NewPlayer", controller.NewPlayer)
		player.POST("/PlayerUpdatePos", controller.PlayerUpdatePos)
	}

	// 登陆和注册
	auth := r.Group("/api/auth")
	{
		// 登陆和注册，不考虑前端所以不需要GET，测试可使用python或postman等
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
		// 用户信息
		auth.GET("/info", AuthMiddleware(), controller.Info)
	}

	lounge := r.Group("/api/lobby")
	{
		lounge.GET("", controller.Info)
		// 进入游戏大厅
		// lounge.GET("", controller.Restroom)
		// 开始匹配
		// lounge.POST("", controller.Match)
	}

	// 都需要加验证中间件
	game := r.Group("/api/game")
	{
		// // 进入游戏
		game.GET("/api/:roomid", controller.Game)
		// // 更新位置
		// game.PATCH("/api/:roomid", controller.Update)
		// // 有人被找到/胜利
		// game.POST("/api/:roomid", controller.Condition)
		// // 退出房间
		// game.DELETE("/api/:roomid", controller.Exit)
		game.GET("", controller.Info)
	}

	Logger().Debug("routes initialized...")

	return r
}
