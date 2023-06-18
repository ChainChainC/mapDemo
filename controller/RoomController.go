package controller

import (
	"mapDemo/common"
	"mapDemo/model"

	"github.com/gin-gonic/gin"
)

func NewRoom(c *gin.Context) {
	req := &model.NewRoomBaseReq{}
	c.Bind(req)
	uuid, err := verifyJwtUuid(req.Jwt)
	if err != nil {
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "NewRoom verifyJwt失败"})
		return
	}
	// TODO:雪花算法产生房间号
	roomId := uuid
	// 查询redis缓存字段，uuid RoomId，如果redis查不到，需要重新登录
	fileds := &[]string{"PlayerType", "RoomId"}
	vals, err := common.LocalRedisClient.GetPlayerInfoByField(uuid, fileds)
	if err != nil {
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "NewRoom 查询用户信息失败"})
		return
	}
	if (*vals)[0] == nil {
		// 玩家不在线
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "NewRoom 玩家缓存丢失，请重新登录"})
		return
	}
	// 判断RoomId是否为空 or 为无效房间号
	if (*vals)[1] == nil || (*vals)[1] == "" {
		// 无效房间号
	} else {
		// 有效房间号，退出之前房间
		rStr, ok := (*vals)[1].(string)
		if !ok {
			c.JSON(200, gin.H{"code": 100, "data": err, "msg": "NewRoom 房间号类型断言失败"})
			return
		}
		if err := common.LocalRedisClient.UpdateRoom(uuid, &rStr, 0); err != nil {
			c.JSON(200, gin.H{"code": 100, "data": err, "msg": "NewRoom 房间退出失败"})
			return
		}
	}
	// 更新玩家信息，不传指针
	err = common.LocalRedisClient.UpdatePlayer(uuid, map[string]interface{}{
		"RoomId":     *roomId,
		"PlayerType": 1,
	})
	if err != nil {
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "PlayerJoinRoom 玩家信息修改失败"})
		return
	}
	// 房间加入玩家uuid
	err = common.LocalRedisClient.UpdateRoom(uuid, roomId, 1)
	if err != nil {
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "PlayerJoinRoom 房间信息更新失败"})
		return
	}
	// 更新位置
	err = common.LocalRedisClient.UpdatePos(uuid, req.Pos)
	if err != nil {
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "PlayerJoinRoom Redis更新pos失败"})
		return
	}
	c.JSON(200, gin.H{"code": 100, "data": req.Pos, "msg": "PlayerJoinRoom 房间信息更新成功"})
}
