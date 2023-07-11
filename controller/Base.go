package controller

import (
	"errors"
	"mapDemo/common"
	"mapDemo/model"
	"math"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

var log = logrus.StandardLogger()

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
		Name:         req.Name,
		Uuid:         req.Uuid,
		PlayerType:   0,
		PlayerPos:    req.PlayerPos,
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

// newJwt分发新的jwt数据
func newJwt(uuid *string) (*string, error) {
	// jwt中存储哪些数据？-->目前只存uuid
	// 分发jwt
	expirationTime := time.Now().Add(2 * time.Hour) // 有效期
	claims := &model.Claims{
		Uuid: *uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),     // 发放时间
			Issuer:    "Map",                 // 发放者
			Subject:   "user token",          // 主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

// 验证jwt并解析uuid
func verifyJwtUuid(jwt *string) (*string, error) {
	if jwt == nil {
		return nil, errors.New("jwt为空")
	}
	// 后续传输的参数可以考虑都通过 Jwt来完成 token 用来干啥
	token, claims, err := common.ParseToken(*jwt)
	if err != nil {
		return nil, err
	}
	// jwt有效
	if token.Valid {
		expirationTime := time.Unix(claims.StandardClaims.ExpiresAt, 0)
		currentTime := time.Now()
		if expirationTime.Before(currentTime) {
			return nil, errors.New("jwt过期")
		}

	} else {
		return nil, errors.New("jwt无效")
	}
	return &claims.Uuid, nil
}

// 计算可见玩家坐标
func getPlayerPosInsight(p *model.Player) []*model.Player {
	allP := model.RoomIdMap[p.RoomId].AllPlayer
	res := make([]*model.Player, 2)
	for k, v := range allP {
		if k != p.Uuid {
			deltaX := p.PlayerPos.X - v.PlayerPos.X
			deltaY := p.PlayerPos.Y - v.PlayerPos.Y
			deltaZ := p.PlayerPos.Z - v.PlayerPos.Z
			distance := math.Sqrt(float64(deltaX*deltaX) + float64(deltaY*deltaY) + float64(deltaZ*deltaZ))
			if distance < float64(p.Sight) {
				res = append(res, v)
			}
		}
	}
	return res
}
