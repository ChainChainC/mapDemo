package common

import (
	"fmt"
	"mapDemo/model"
	"testing"
)

func TestReleaseToken(t *testing.T) {
	var player model.Player
	player.Name = "cc"
	player.Uuid = "123456"
	res, err := ReleaseToken(player)
	if err != nil {
		t.Fail()
	}
	fmt.Println(res)
}

func TestParseToken(t *testing.T) {
	var player model.Player
	player.Name = "cc"
	player.Uuid = "123456"
	res, err := ReleaseToken(player)
	if err != nil {
		t.Fail()
	}
	// fmt.Println(res)
	token, clams, err := ParseToken(res)
	fmt.Println(token)
	if clams.PlayerUuid != player.Uuid {
		t.Error("解析玩家uuid错误")
	}
	fmt.Println(clams.PlayerUuid)
}
