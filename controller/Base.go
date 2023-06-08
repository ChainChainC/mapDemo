package controller

import (
	"errors"
	"mapDemo/common"
	"mapDemo/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 新玩家连入
func newPlayer(req *model.NewPlayerReq) *model.Player {
	// 分发jwt
	expirationTime := time.Now().Add(2 * time.Hour) // 有效期
	claims := &model.Claims{
		Uuid: req.Uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),     // 发放时间
			Issuer:    "binbin",              // 发放者
			Subject:   "user token",          // 主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)
	np := &model.Player{
		Name:       "test player",
		Uuid:       req.Uuid,
		PlayerType: 0,
		PlayerPos: model.Pos{
			X: 1.0,
			Y: 1.0,
			Z: 1.0,
		},
		RoomId:       "",
		InRoom:       false,
		PlayerOnline: true,
		PlayerJwt:    tokenString,
	}
	// 创建好的玩家加入到全局玩家Map中
	// TODO：并发性问题
	model.PlayerIdMap[np.Uuid] = np
	return np
}

// 验证Jwt
func verifyJwt(p *model.Player, jwt *string) error {
	// 后续传输的参数可以考虑都通过 Jwt来完成 token 用来干啥
	token, claims, err := common.ParseToken(*jwt)
	if err != nil {
		return err
	}
	// jwt有效
	if token.Valid {
		expirationTime := time.Unix(claims.StandardClaims.ExpiresAt, 0)
		currentTime := time.Now()
		if expirationTime.Before(currentTime) {
			return errors.New("jwt过期")
		}

	} else {
		return errors.New("jwt无效")
	}
	return nil
}

// 玩家加入房间
func addPlayerIntoRoom(p *model.Player, roomId *model.IdentifyType) error {
	// 如果玩家已经在房间
	if p.InRoom {
		deletePlayerFromRoom(p)
	}
	// 判断房间是否存在
	if r, ok := model.RoomIdMap[*roomId]; ok {
		// 玩家加入房间列表列表
		r.AllPlayer[p.Uuid] = p
		p.InRoom = true
		p.RoomId = *roomId
	} else {
		// 房间不存在
		// 统一Error处理
		return errors.New("加入的房间不存在")
	}
	return nil
}

// 将玩家从房间剔除逻辑
func deletePlayerFromRoom(p *model.Player) error {
	if p.InRoom {

	} else {
		return errors.New("玩家不在房间内")
	}
	// 定位房间
	// 需要上锁
	delete(model.RoomIdMap[p.RoomId].AllPlayer, p.Uuid)

	return nil
}

// 计算可见玩家坐标
func getPlayerPosInsight() {

}
