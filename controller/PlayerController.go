package controller

import (
	"mapDemo/model"

	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("a_secret_crect")

func NewPlayer(c *gin.Context) {
	npReq := &model.NewPlayerReq{}
	c.BindJSON(npReq)
	// 去全局中查找玩家是否存在
	if p, ok := model.PlayerIdMap[npReq.Uuid]; ok {
		// 玩家存在，判断jwt是否过期
		if p.Uuid == "" {
			// TODO: 抛出异常
		}
		if p.PlayerJwt == "" {
			// TODO: 分发jwt
		}
		//
		c.JSON(200, gin.H{"code": 200, "data": "v", "msg": "玩家重连接入"})
	} else {
		// TODO:如果发现openId为空, 抛出异常
		np := newPlayer(npReq)
		c.JSON(200, gin.H{"code": 200, "data": np.Uuid + " Jwt: " + np.PlayerJwt, "msg": "新玩家接入"})
	}

}

// 玩家后续隔一段时间向服务器发起更新位置请求
func PlayerUpdatePos(c *gin.Context) {
	pupReq := &model.PlayerUpdatePosReq{}
	c.BindJSON(pupReq)
	// 获取玩家数据
	// 更新玩家位置
	if p, ok := model.PlayerIdMap[pupReq.Uuid]; ok {
		// TODO：做玩家合法性验证, 验证玩家各项信息合法性（是否超时）
		verifyJwt(p, &pupReq.Jwt)

		// 更新玩家位置
		p.PlayerPos = pupReq.PlayerPos
		// 玩家在房间内，那么需要走另外的处理逻辑
		if p.InRoom {
			// 返回所有可见玩家位置
			c.JSON(200, gin.H{"code": 200, "data": "其它可见玩家坐标", "msg": "玩家位置更新成功，返回可见玩家位置"})
		} else {
			c.JSON(200, gin.H{"code": 200, "data": p.PlayerPos, "msg": "玩家位置更新成功"})
		}
	} else {
		// TODO: 玩家状态不在线，抛出提示,让玩家重新登录
		c.JSON(200, gin.H{"code": 100, "data": nil, "msg": "玩家不在服务器，请重新登录"})
	}
}

func PlayerJoinRoom(c *gin.Context) {
	pjrReq := &model.PlayerJoinRoomReq{}
	c.BindJSON(pjrReq)
	if v, ok := model.PlayerIdMap[pjrReq.Uuid]; ok {
		if r, ok := model.RoomIdMap[pjrReq.RoomUuid]; ok {
			// 玩家加入列表
			r.AllPlayer[v.Uuid] = v
			// TODO:返回其他玩家位置 做成工具, 会频繁调用
			c.JSON(200, gin.H{"code": 100, "data": "返回房间内其它玩家位置", "msg": "玩家成功加入房间"})
		} else {
			// 报错, 房间号不存在
			c.JSON(100, gin.H{"code": 100, "data": nil, "msg": "加入房间失败, 房间不存在"})
		}
	} else {
		// TODO: 处理玩家不在线
		c.JSON(100, gin.H{"code": 100, "data": nil, "msg": "玩家加入房间，玩家状态不在线"})
	}
}

func PlayerQuitRoom(c *gin.Context) {
	pqrReq := &model.PlayerQuitRoomReq{}
	c.BindJSON(pqrReq)
	// 获取玩家对象
	if p, ok := model.PlayerIdMap[pqrReq.Uuid]; ok {
		// TODO: 合法性验证
		verifyJwt(p, &pqrReq.Jwt)
		if p.InRoom {
			deletePlayerFromRoom(p)
		} else {
			c.JSON(200, gin.H{"code": 100, "data": "clams", "msg": "玩家不在房间内"})
		}
	} else {
		// TODO:玩家未在服务器内登录记录
		c.JSON(100, gin.H{"code": 100, "data": "退出房间操作", "msg": "玩家不在线"})
	}
}
