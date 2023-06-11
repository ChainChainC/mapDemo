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

func TestBase(t *testing.T) {
	NewRedisClientApp()
	defer LocalRedisClient.Client.Close()
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
	if err := LocalRedisClient.Client.Set(ctx, key, valStr, 120*time.Second).Err(); err != nil {
		fmt.Println("写入错误")
		t.Fail()
	}
	// 获取数据
	if result, err := LocalRedisClient.Client.Get(ctx, key).Result(); err != nil {
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
	defer LocalRedisClient.Client.Close()
	// 需要在初始化时检测是否连接成功
	ctx := context.Background()
	key := "Room:roomUuid"
	val := "playerUuid1"
	val1 := []string{val, "Uuid2", "Uuid3"}
	// 这里key 也可以为 []string类型
	if err := LocalRedisClient.Client.SAdd(ctx, key, val1).Err(); err != nil {
		fmt.Printf("Room加入玩家错误")
		t.Fail()
	} else {
		fmt.Println("元素加入成功")
	}
	// 获取set当中的所有元素
	if smembers, err := LocalRedisClient.Client.SMembers(ctx, key).Result(); err != nil {
		fmt.Printf("获取set元素错误")
		t.Fail()
	} else {
		// []string
		fmt.Println(reflect.TypeOf(smembers))
		fmt.Println(smembers)
	}
	// 获取set中元素的数量
	if scards, err := LocalRedisClient.Client.SCard(ctx, key).Result(); err != nil {
		fmt.Printf("获取set元素错误")
		t.Fail()
	} else {
		fmt.Printf("数量: %d", scards)
	}
	// SPOP，随机移除一个数据并返回这个数据
	// SISMEMBER，判断元素是否在集合中
	// 以及交集并集等操作
	// SREM , 删除值，返回删除元素个数
	if count, err := LocalRedisClient.Client.SRem(ctx, key, val).Result(); err != nil {
		fmt.Printf("删除元素错误")
		t.Fail()
	} else {
		fmt.Printf("删除数量：%d", count)
	}
}
