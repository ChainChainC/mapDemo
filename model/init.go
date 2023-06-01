package model

// -----定义model内基本类型----
// 定义M类型
type M map[string]interface{}

// 坐标
type Pos struct {
	x float32
	y float32
	z float32
}

// 存储所有房间的全局MAP
// TODO, 后续全局Map考虑放入redis, string or int ->于paler下的roomId关联
var RoomIdMap = make(M, 4)

// func Init() {
// 	RoomIdMap := make(M, 10)
// }
