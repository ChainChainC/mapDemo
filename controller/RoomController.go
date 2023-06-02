package controller

import (
	"mapDemo/model"

	"github.com/gin-gonic/gin"
)

func NewRoom(ctx *gin.Context) {
	req := &model.NewRoomReq{}
	ctx.Bind(req)
	r := &model.Room{
		RoomId:    req.Uuid,
		AllPlayer: make(map[model.IdentifyType]*model.Player, 4),
	}
	// 创建好的房间加入全局表
	model.RoomIdMap[r.RoomId] = r
	ctx.JSON(200, gin.H{
		"message": "房间创建成功" + r.RoomId,
	})
	// return r
}
