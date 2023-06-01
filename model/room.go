package model

type Room struct {
	// 可以不存?
	RoomId string
	// 存储所有玩家指针(切片)
	AllPlayer []*Player
	// TODO Map信息
}

func NewRoom() *Room {
	r := &Room{
		RoomId:    "test",
		AllPlayer: make([]*Player, 4),
	}
	// 创建好的房间加入全局表
	RoomIdMap[r.RoomId] = r
	return r
}
