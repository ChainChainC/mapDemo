package common

import (
	"context"
	"encoding/json"
	"fmt"
	"mapDemo/model"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClientApp struct {
	Client *redis.Client
}

var LocalRedisClient *RedisClientApp

func NewRedisClientApp() {
	LocalRedisClient = &RedisClientApp{
		Client: redis.NewClient(
			&redis.Options{
				Addr:     "localhost:6379",
				Password: "12345",
				DB:       10,
			}),
	}
	// 检查连接redis是否成功
	if err := LocalRedisClient.checkClient(); err != nil {
		panic(err)
	}
}

// 更新坐标到redis
func (c *RedisClientApp) UpdatePos(pUuid *string, pos *model.Pos) error {
	ctx := context.Background()
	// TODO：加入分布式锁
	valStr, err := json.Marshal(pos)
	if err != nil {
		fmt.Printf("struct2str 错误")
		return err
	}
	// 字符串形式写入，超时时间为120s
	if err := c.Client.Set(ctx, *pUuid, valStr, 120*time.Second).Err(); err != nil {
		fmt.Println("写入错误")
		return err
	}
	return nil
}

func (c *RedisClientApp) UpdatePlayer(p *model.Player) {

}

// 更新房间中玩家信息、Set 操作
// TODO: 后续需要开脚本定期清除没有玩家的房间
// action 0 删除 1加入
func (c *RedisClientApp) UpdateRoom(pUuid *string, rId *string, action uint8) error {
	ctx := context.Background()
	if action == 0 {
		if _, err := c.Client.SRem(ctx, "Room:"+*rId, *pUuid).Result(); err != nil {
			return err
		}
	} else if action == 1 {
		if err := c.Client.SAdd(ctx, "Room:"+*rId, *pUuid).Err(); err != nil {
			return err
		}
	}
	return nil
}

// 检查连接redis是否成功
func (c *RedisClientApp) checkClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if _, err := c.Client.Ping(ctx).Result(); err != nil {
		fmt.Printf("Connect Failed: %v \n", err)
		return err
	}
	return nil
}
