package model

type Player struct {
	// ---玩家信息
	Name string
	// uuid 或者其它能够标识玩家唯一性数据
	// TODO: string 或 int待定
	Uuid string
	// 玩家身份
	PlayerType int8
	// 玩家坐标
	PlayerPos Pos
	// -----房间信息
	// 房间号 和玩家uuid类似, string 或 int
	RoomId string
	// 玩家是否在房间内
	InRoom bool
}

// 初始化一个玩家, 返回指针
func NewPlayer() *Player {
	// np := &Player{
	// 	RoomId: "test",
	// }
	// RoomIdMap[np.RoomId] = append(RoomIdMap[np.RoomId], np)
	return nil
}
