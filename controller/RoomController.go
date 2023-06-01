package controller

import (
	"mapDemo/model"

	"github.com/gin-gonic/gin"
)

func NewRoom(ctx *gin.Context) {
	r := model.NewRoom()
	ctx.JSON(200, gin.H{
		"message": "房间创建成功" + r.RoomId,
	})
}
