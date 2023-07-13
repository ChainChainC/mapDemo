package common

import (
	"context"
	"encoding/json"
	"mapDemo/model"
	"time"

	"github.com/go-redis/redis/v8"
)

var LocalRedisClient *RedisClientApp

const (
	playerPrefix = "Player:"
	roomPrefix   = "Room:"
	posPrefix    = "Pos:"
)

type RedisClientApp struct {
	client *redis.Client
}

// 检查连接redis是否成功
func (c *RedisClientApp) checkClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if _, err := c.client.Ping(ctx).Result(); err != nil {
		// fmt.Printf("Connect Failed: %v \n", err)
		log.WithError(err).Error("Redis connect FAILED.")
		return err
	}
	return nil
}

func NewRedisClientApp() {
	LocalRedisClient = &RedisClientApp{
		client: redis.NewClient(
			&redis.Options{
				Addr: "localhost:6379",
				// Password: "12345",
				DB: 10,
			}),
	}
	// 检查连接redis是否成功
	if err := LocalRedisClient.checkClient(); err != nil {
		panic(err)
	}
}

// UpdatePos 更新坐标到redis
func (c *RedisClientApp) UpdatePos(pUuid *string, pos *model.Pos) error {
	ctx := context.Background()
	// TODO：加入分布式锁
	valStr, err := json.Marshal(pos)
	if err != nil {
		// fmt.Printf("struct2str 错误")
		log.WithError(err).Error("Redis struct2str 错误.")
		return err
	}
	// 字符串形式写入，超时时间为120s
	if err := c.client.Set(ctx, posPrefix+*pUuid, valStr, 120*time.Minute).Err(); err != nil {
		// fmt.Println("写入错误")
		log.WithError(err).Error("Redis 写入错误.")
		return err
	}
	return nil
}

// UpdatePlayer 玩家信息更新，any不可以为指针或含有指针类型
func (c *RedisClientApp) UpdatePlayer(key *string, val any) error {
	ctx := context.Background()
	if err := c.client.HMSet(ctx, playerPrefix+*key, val).Err(); err != nil {
		log.WithError(err).Error("Redis 数据插入错误.")
		return err
	} else {
		return nil
	}
}

// UpdateRoom 更新房间中玩家信息、Set 操作
// TODO: 后续需要开脚本定期清除没有玩家的房间
// action 0 删除 1加入
func (c *RedisClientApp) UpdateRoom(pUuid *string, rId *string, action uint8) error {
	ctx := context.Background()
	if action == 0 {
		if _, err := c.client.SRem(ctx, roomPrefix+*rId, *pUuid).Result(); err != nil {
			return err
		}
	} else if action == 1 {
		if err := c.client.SAdd(ctx, roomPrefix+*rId, *pUuid).Err(); err != nil {
			return err
		}
	}
	return nil
}

// GetPlayerInfoByField 获取用户详细信息 Hash结果对象
func (c *RedisClientApp) GetPlayerInfoByField(key *string, fields *[]string) (*[]interface{}, error) {
	ctx := context.Background()
	vals, err := LocalRedisClient.client.HMGet(ctx, playerPrefix+*key, *fields...).Result()
	if err != nil {
		return nil, err
	}
	return &vals, nil
}

// IsPlayerOnline 快速查询玩家是否在线（是否在redis中）
func (c *RedisClientApp) IsPlayerOnline(key *string) (bool, error) {
	ctx := context.Background()
	exist, err := c.client.Exists(ctx, playerPrefix+(*key)).Result()
	if err != nil {
		return false, err
	} else {
		return exist == 1, nil
	}
}

func (c *RedisClientApp) DeletePlayerInfo(key *string) (int64, error) {
	ctx := context.Background()
	count, err := c.client.Del(ctx, playerPrefix+(*key)).Result()
	if err != nil {
		return count, err
	}
	return count, nil
}
