package controller

import (
	"mapDemo/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 新玩家连入
func newPlayer(req *model.NewPlayerReq) *model.Player {
	// 分发jwt
	expirationTime := time.Now().Add(24 * time.Hour) // 有效期
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

func verifyJwt(p *model.Player, jwt *string) {
	// 验证jwt
	// token, clams, err := common.ParseToken(p.PlayerJwt)
	// if err != nil {
	// 	c.JSON(100, gin.H{"code": 100, "data": "退出房间操作", "msg": "jwt parase 失败"})
	// }
}

// 将玩家从房间剔除逻辑
func deletePlayerFromRoom(p *model.Player) {
	// 定位房间
	// 需要上锁
	delete(model.RoomIdMap[p.RoomId].AllPlayer, p.Uuid)
}
