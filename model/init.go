package model

import "github.com/dgrijalva/jwt-go"

// -----定义model内基本类型----
// 定义M类型
type (
	M            map[IdentifyType]interface{}
	IdentifyType = string // 玩家 & 房间的唯一标识类型
)

// 坐标
type Pos struct {
	X float32
	Y float32
	Z float32
}

// 存储所有房间的全局MAP
// TODO, 后续全局Map考虑放入redis, string or int ->于paler下的roomId关联
var (
	RoomIdMap   = make(map[IdentifyType]*Room, 4)
	PlayerIdMap = make(map[IdentifyType]*Player, 32)
)

type Claims struct {
	Uuid IdentifyType
	jwt.StandardClaims
}
