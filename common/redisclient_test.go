package common

import (
	"context"
	"encoding/json"
	"fmt"
	"mapDemo/model"
	"reflect"
	"testing"
	"time"
)

// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)可以设置redis操作的超时时间

func TestBase(t *testing.T) {
	NewRedisClientApp()
	defer LocalRedisClient.client.Close()
	ctx := context.Background()
	key := "Pos:ccbb"
	value := model.Pos{
		X: 1,
		Y: 2,
		Z: 0,
	}
	valStr, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("struct2str 错误")
	}
	// 字符串形式写入
	if err := LocalRedisClient.client.Set(ctx, key, valStr, 120*time.Second).Err(); err != nil {
		fmt.Println("写入错误")
		t.Fail()
	}
	// 获取数据
	if result, err := LocalRedisClient.client.Get(ctx, key).Result(); err != nil {
		t.Fail()
	} else {
		var pos model.Pos
		if err := json.Unmarshal([]byte(result), &pos); err != nil {
			fmt.Println("解析错误")
			t.Fail()
		}
		fmt.Println(result)
		fmt.Println(pos.X)
	}
}

func TestSet(t *testing.T) {
	NewRedisClientApp()
	defer LocalRedisClient.client.Close()
	// 需要在初始化时检测是否连接成功
	ctx := context.Background()
	key := "Room:roomUuid"
	val := "playerUuid1"
	val1 := []string{val, "Uuid2", "Uuid3"}
	// 这里key 也可以为 []string类型
	if err := LocalRedisClient.client.SAdd(ctx, key, val1).Err(); err != nil {
		fmt.Printf("Room加入玩家错误")
		t.Fail()
	} else {
		fmt.Println("元素加入成功")
	}
	// 获取set当中的所有元素
	if smembers, err := LocalRedisClient.client.SMembers(ctx, key).Result(); err != nil {
		fmt.Printf("获取set元素错误")
		t.Fail()
	} else {
		// []string
		fmt.Println(reflect.TypeOf(smembers))
		fmt.Println(smembers)
	}
	// 获取set中元素的数量
	if scards, err := LocalRedisClient.client.SCard(ctx, key).Result(); err != nil {
		fmt.Printf("获取set元素错误")
		t.Fail()
	} else {
		fmt.Printf("数量: %d", scards)
	}
	// SPOP，随机移除一个数据并返回这个数据
	// SISMEMBER，判断元素是否在集合中
	// 以及交集并集等操作
	// SREM , 删除值，返回删除元素个数
	if count, err := LocalRedisClient.client.SRem(ctx, key, val).Result(); err != nil {
		fmt.Printf("删除元素错误")
		t.Fail()
	} else {
		fmt.Printf("删除数量：%d", count)
	}
}

type HashUpdate struct {
	PlayerType int8
	RoomId     string
	Sight      int
}

func TestHash(t *testing.T) {
	NewRedisClientApp()
	defer LocalRedisClient.client.Close()
	// 用户更新key: Player:uuid map[string]interface{}
	key := "Player:uuid1"
	// 如果直接结构体，插入会报错
	// val := &HashUpdate{
	// 	PlayerType: 1,
	// 	RoomId:     "Aaaaaaaa",
	// 	Sight:      100,
	// }
	val := map[string]interface{}{
		"PlayerType": 1,
		"RoomId":     "Aaaaaaaa",
		"Sight":      100,
	}
	var a any
	a = val
	// 需要在初始化时检测是否连接成功
	ctx := context.Background()
	// HSet插入数据--->只能单个 key val插入
	// HMSet，使用map进行插入
	if err := LocalRedisClient.client.HMSet(ctx, key, a).Err(); err != nil {
		fmt.Printf("数据插入错误")
		t.Fail()
	} else {

		fmt.Printf("数据写入成功")
	}
	// HMGET, 根据hash key和多个字段获取值
	if vals, err := LocalRedisClient.client.HMGet(ctx, "key", "RoomId", "Sight").Result(); err != nil {
		fmt.Printf("数据获取错误")
		t.Fail()
	} else {
		fmt.Println(reflect.TypeOf(vals))
		fmt.Println(vals[0])
		fmt.Print(vals[0] == nil)
	}
	// HLen，获取hashkey的字段多少
	// HDel，删除字段，支持删除多个字段

	// 设置键的过期时间为 120 秒
	err := LocalRedisClient.client.Expire(ctx, key, 120*time.Second).Err()
	if err != nil {
		fmt.Println("Failed to set expiration:", err)
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	NewRedisClientApp()
	defer LocalRedisClient.client.Close()
	// 用户更新key: Player:uuid map[string]interface{}
	key := "uuid1"
	err := LocalRedisClient.UpdatePlayer(&key, map[string]interface{}{
		"PlayerType": 0,
		"RoomId":     "",
	})
	if err != nil {
		fmt.Printf("数据写入错误")
	}
	count, err := LocalRedisClient.DeletePlayerInfo(&key)
	if err != nil {
		fmt.Printf("数据删除错误")
		t.Fail()
	} else {
		fmt.Printf("数据删除成功: %d", count)
	}
}
