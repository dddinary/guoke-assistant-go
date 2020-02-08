package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"guoke-helper-golang/config"
	"log"
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

func ValidateTokenByRedis(token string) int {
	res, err := RedisCli.Get(token).Int()
	if err == redis.Nil {
		log.Printf("redis中token不存在")
		return 0
	} else if err != nil {
		log.Printf("redis中token获取uid出错：%v\n", err)
		return 0
	}
	return res
}

func AddTokenToRedis(token string, uid int) error {
	err := RedisCli.Set(token, uid, 0).Err()
	if err != nil {
		log.Printf("redis中token设置uid出错：%v\n", err)
	}
	return nil
}

func DeleteTokenInRedis(token string) error {
	err := RedisCli.Del(token).Err()
	if err != nil {
		log.Printf("redis中token删除uid出错：%v\n", err)
		return err
	}
	return nil
}
