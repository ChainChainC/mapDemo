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
	token, clamis, err := ParseToken(res)
	fmt.Println(token)
	if clamis.Uuid != player.Uuid {
		t.Error("解析玩家uuid错误")
	}
	fmt.Println(clamis.Uuid)
	fmt.Println(clamis.StandardClaims.ExpiresAt)
}
