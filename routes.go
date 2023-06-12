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
		player.POST("/PlayerJoinRoom", controller.PlayerJoinRoom)
		player.POST("/PlayerQuitRoom", controller.PlayerQuitRoom)
	}

	// // 都需要加验证中间件
	// game := r.Group("/api/game")
	// {
	// 	// // 进入游戏
	// 	game.GET("/api/:roomid", controller.Game)
	// 	// // 更新位置
	// 	// game.PATCH("/api/:roomid", controller.Update)
	// 	// // 有人被找到/胜利
	// 	// game.POST("/api/:roomid", controller.Condition)
	// 	// // 退出房间
	// 	// game.DELETE("/api/:roomid", controller.Exit)
	// 	game.GET("", controller.Info)
	// }

	Logger().Debug("routes initialized...")

	return r
}
