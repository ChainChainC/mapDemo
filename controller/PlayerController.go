package controller

import (
	"fmt"
	"mapDemo/common"
	"mapDemo/model"

	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("a_secret_crect")

func NewPlayer(c *gin.Context) {
	newPlayerRedis(c)
}

func newPlayerRedis(c *gin.Context) {
	baseReq := &model.NewPlayerBaseReq{}
	c.BindJSON(baseReq)
	if baseReq.Jwt != nil {
		// 解析jwt，写入redis
		uuid, err := verifyJwtUuid(baseReq.Jwt)
		if err != nil {
			// TODO：jwt验证过期另外逻辑处理
			c.JSON(200, gin.H{"code": 100, "data": err, "msg": "VerifyJwt失败"})
			log.Error("Player jwt verify FAILED.")
			return
		}
		exist, err := common.LocalRedisClient.IsPlayerOnline(uuid)
		if err != nil {
			log.Error("NewPlayer exist verify with redis FAILED.")
			return
		}
		// 在redis中查询到信息
		if exist {
			c.JSON(200, gin.H{"code": 200, "data": *baseReq.Code, "msg": "玩家存在"})
			return
		} else {
			// TODO：完善
			_, err := common.LocalRedisClient.DeletePlayerInfo(uuid)
			if err != nil {
				log.WithError(err).Error("del info with redis FAILED.")
				return
			}
			c.JSON(200, gin.H{"code": 200, "data": *baseReq.Code, "msg": "清除之前缓存请重新登陆"})
		}
	} else { // 调用接口且不携带jwt
		if baseReq.Code == nil {
			c.JSON(200, gin.H{"code": 200, "data": "", "msg": "新玩家接入，未传入code"})
			log.Warn("Player connect without Code.")
			return
		}
		// 用code获取openId
		openId := baseReq.Code
		// 签发jwt
		jwt, err := newJwt(openId)
		if err != nil {
			c.JSON(200, gin.H{"code": 200, "data": *baseReq.Code, "msg": "新玩家接入时jwt签发失败"})
			log.Error("New Player con with jwt FAILED.")
			return
		}
		// TODO：直接更新，如果玩家在线，是不是无法连接回房间内
		// 不能传入指针
		err = common.LocalRedisClient.UpdatePlayer(openId, map[string]interface{}{
			"PlayerType": 0,
			"RoomId":     "",
			// "PlayerOnline": 1,
		})
		if err != nil {
			c.JSON(200, gin.H{"code": 200, "data": *baseReq.Code, "msg": "玩家信息写入redis失败"})
			log.Error("New player redis write FAILED.")
		}
		c.JSON(200, gin.H{"code": 200, "data": *baseReq.Code + " Jwt: " + *jwt, "msg": "新玩家接入"})
	}
}

// 玩家后续隔一段时间向服务器发起更新位置请求
func PlayerUpdatePos(c *gin.Context) {
	req := &model.PlayerUpdatePosBaseReq{}
	c.BindJSON(req)
	uuid, err := verifyJwtUuid(req.Jwt)
	if err != nil {
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "updatePos verifyJwt失败"})
		log.Error("Player jwt verify FAILED.")
		return
	}
	// 更新Pos
	err = common.LocalRedisClient.UpdatePos(uuid, req.Pos)
	if err != nil {
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "updatePos Redis更新pos失败"})
		log.Error("Player pos update with redis FAILED.")
		return
	}
	// 更新缓存
	updatePlayerPosCache(uuid, req.Pos)
	// 在房间内
	// TODO：是否在房间内需要从前端携带Type？那如果玩家掉线前端缓存出错会不会产生问题
	if req.Type != 0 {
		// TODO：获取全部玩家坐标，并判断可见性
	} else {
		c.JSON(200, gin.H{"code": 100, "data": req.Pos, "msg": "更新玩家坐标成功，玩家不在房间内"})
		return
	}
}

// PlayerJoinRoom 玩家申请加入房间
func PlayerJoinRoom(c *gin.Context) {
	req := &model.PlayerJoinRoomBaseReq{}
	if err := c.BindJSON(req); err != nil {
		return
	}
	uuid, err := verifyJwtUuid(req.Jwt)
	if err != nil {
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "PlayerJoinRoom verifyJwt失败"})
		log.Error("Player jwt verified FAILED.")
		return
	}
	// 查询redis缓存字段，uuid RoomId，如果redis查不到，需要重新登录
	fileds := &[]string{"PlayerType", "RoomId"}
	vals, err := common.LocalRedisClient.GetPlayerInfoByField(uuid, fileds)
	if err != nil || vals == nil {
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "PlayerJoinRoom 查询用户信息失败"})
		log.Error("Player join room with redis FAILED.")
		return
	}
	fmt.Print(vals)
	if (*vals)[0] == nil {
		// 玩家不在线
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "PlayerJoinRoom 玩家缓存丢失，请重新登录"})
		return
	}
	// 判断RoomId是否为空 or 为无效房间号
	if (*vals)[1] == nil || (*vals)[1] == "" {
		// 无效房间号
	} else {
		// 有效房间号，退出之前房间
		rStr, ok := (*vals)[1].(string)
		if !ok {
			c.JSON(200, gin.H{"code": 100, "data": err, "msg": "PlayerJoinRoom 房间号类型断言失败"})
			log.Error("Player room id judge FAILED.")
			return
		}
		if err := common.LocalRedisClient.UpdateRoom(uuid, &rStr, 0); err != nil {
			c.JSON(200, gin.H{"code": 100, "data": err, "msg": "PlayerJoinRoom 房间退出失败"})
			log.Error("Player quit room with redis FAILED.")
			return
		}
	}
	// 更新玩家信息
	err = common.LocalRedisClient.UpdatePlayer(uuid, map[string]interface{}{
		"RoomId":     *req.RoomId,
		"PlayerType": 1,
	})
	if err != nil {
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "PlayerJoinRoom 玩家信息修改失败"})
		log.Error("Player info update with redis FAILED.")
		return
	}
	// 房间加入玩家uuid
	err = common.LocalRedisClient.UpdateRoom(uuid, req.RoomId, 1)
	if err != nil {
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "PlayerJoinRoom 房间信息更新失败"})
		log.Error("Player room info update with redis FAILED.")
		return
	}
	// TODO：缓存房间内玩家
	// TODO：获取房间内玩家坐标返回
}

// PlayerQuitRoom 玩家退出房间
func PlayerQuitRoom(c *gin.Context) {
	req := &model.PlayerQuitRoomBaseReq{}
	c.BindJSON(req)
	uuid, err := verifyJwtUuid(req.Jwt)
	if err != nil {
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "PlayerQuitRoom verifyJwt失败"})
		log.Error("Player jwt verify FAILED.")
		return
	}
	if err := common.LocalRedisClient.UpdateRoom(uuid, req.RoomId, 0); err != nil {
		c.JSON(200, gin.H{"code": 100, "data": err, "msg": "PlayerQuitRoom 房间退出失败"})
		log.Error("Player quit room with redis FAILED.")
		return
	}
	err = common.LocalRedisClient.UpdatePlayer(uuid, map[string]interface{}{
		"RoomId":     "",
		"PlayerType": 0,
	})
	if err != nil {
		log.Error("Player quit room update info with redis FAILED.")
		return
	}
	// TODO：房间内玩家缓存修改
}
