package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"guoke-helper-golang/config"
)

var RedisCli *redis.Client

func init() {
	redisHost := config.RedisConf.Host
	redisPort := config.RedisConf.Port
	redisAuth := config.RedisConf.AUTH
	RedisCli = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password:           redisAuth,
		DB:                 0,
	})
}

