package model

type Room struct {
	// 可以不存?
	RoomId IdentifyType
	// 存储所有玩家指针(切片)-->玩家退出的话，需要及时判断退出玩家，并将其从切片中删除
	// TODO：删除操作要上锁，可能存在并发问题 -> 读写锁
	AllPlayer map[IdentifyType]*Player
	// room 状态
	RoomState int8
	// TODO：Map信息
}

// NewRoomReq ------------- Redis version ------------
type NewRoomBaseReq struct {
	Jwt *string `json:"jwt"`
	Pos *Pos    `json:"pos"`
}
