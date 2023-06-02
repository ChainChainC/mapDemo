package controller

import (
	"mapDemo/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("a_secret_crect")

func NewPlayer(c *gin.Context) {
	npReq := &model.NewPlayerReq{}
	c.BindJSON(npReq)
	// 去全局中查找玩家是否存在
	if v, ok := model.PlayerIdMap[npReq.Uuid]; ok {
		// 玩家存在，判断jwt是否过期
		if v.Uuid == "" {
			// TODO: 抛出异常
		}
		if v.PlayerJwt == "" {
			// TODO: 分发jwt
		}
		//
		c.JSON(200, gin.H{"code": 200, "data": "v", "msg": "玩家重连接入"})
	} else {
		// TODO:如果发现openId为空, 抛出异常
		expirationTime := time.Now().Add(24 * time.Hour) // 有效期
		claims := &model.Claims{
			Uuid: npReq.Uuid,
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
			Name:         "test player",
			Uuid:         npReq.Uuid,
			PlayerType:   0,
			PlayerPos:    model.Pos{},
			RoomId:       "",
			InRoom:       false,
			PlayerOnline: true,
			PlayerJwt:    tokenString,
		}
		// 创建好的玩家加入到全局玩家Map中
		model.PlayerIdMap[np.Uuid] = np
		c.JSON(200, gin.H{"code": 200, "data": np.Uuid + " Jwt: " + np.PlayerJwt, "msg": "新玩家接入"})
	}

}

// 玩家连接进入服务器
func PlayerUpdatePos(c *gin.Context) {

}
