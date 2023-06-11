package common

import "github.com/go-redis/redis"

type RedisClientApp struct {
	Client *redis.Client
}

var LocalRedisClient *RedisClientApp

func NewRedisClientApp() {
	LocalRedisClient = &RedisClientApp{
		Client: redis.NewClient(
			&redis.Options{
				Addr:     "localhost:6379",
				Password: "",
				DB:       10,
			}),
	}
}
