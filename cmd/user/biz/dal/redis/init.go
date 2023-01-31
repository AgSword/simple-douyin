package redis

import (
	"context"

	"github.com/AgSword/simple-douyin/cmd/user/conf"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
)

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
