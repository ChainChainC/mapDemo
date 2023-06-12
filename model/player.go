package model

type Player struct {
	// 玩家身份
	PlayerType int8
	// 房间号 和玩家uuid类似, string 或 int
	RoomId IdentifyType
	// 玩家视野
	Sight uint32
	// ---玩家信息
	Name string
	// uuid 或者其它能够标识玩家唯一性数据
	// TODO: string 或 int待定
	Uuid IdentifyType
	// 玩家坐标
	PlayerPos Pos
	// -----房间信息

	// 玩家是否在房间内
	InRoom bool
	// TODO，玩家在线状态，下线一段时间后需要从Map中清除玩家
	PlayerOnline bool
	// 玩家token
	PlayerJwt string
}

// NewPlayerReq 请求体
type NewPlayerReq struct {
	// ---玩家信息
	Name      string       `json:"nickName"`
	Uuid      IdentifyType `json:"openId"`
	PlayerPos Pos          `json:"playerPos"`
}

//

// 玩家更新坐标
type PlayerUpdatePosReq struct {
	Name      string       `json:"nickName"`
	Uuid      IdentifyType `json:"openId"`
	PlayerPos Pos          `json:"playerPos"`
	Jwt       string       `json:"jwt"`
}

// 玩家加入房间
type PlayerJoinRoomReq struct {
	Name     string       `json:"nickName"`
	Uuid     IdentifyType `json:"openId"`
	RoomUuid IdentifyType `json:"roomUuid"`
	Jwt      string       `json:"jwt"`
}

type PlayerQuitRoomReq struct {
	Name     string       `json:"nickName"`
	Uuid     IdentifyType `json:"openId"`
	RoomUuid IdentifyType `json:"roomUuid"`
	Jwt      string       `json:"jwt"`
}

// ---------------Redis精简--------------

type PlayerBaseReq struct {
	Jwt string `json:"jwt"`
}

type NewPlayerBaseReq struct {
	Code *string `Json:"code"`
	Jwt  *string `json:"jwt"`
}

type PlayerUpdatePosBaseReq struct {
	Jwt    *string `json:"jwt"`
	Pos    *Pos    `json:"pos"`
	Type   uint8   `json:"type"`
	RoomId *string `json:"roomId"`
}

type PlayerJoinRoomBaseReq struct {
	Jwt    *string `json:"jwt"`
	RoomId *string `json:"roomId"`
	Pos    *Pos    `json:"pos"`
}

type PlayerQuitRoomBaseReq struct {
	Jwt    *string `json:"jwt"`
	RoomId *string `json:"roomId"`
}
