package service

import (
	"github.com/go-redis/redis"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/utils"
	"log"
	"time"
)

var BitCountRange = &redis.BitCount{Start: 0, End: 100000}


func SetUserDAU(uid int) {
	dayStr := time.Now().Format("20060102")
	auKey := constant.RedisKeyPrefixAU + dayStr
	_ = utils.RedisCli.SetBit(auKey, int64(uid), 1).Err()
}

func SetIpDUV(ip string) {
	dayStr := time.Now().Format("20060102")
	uvKey := constant.RedisKeyPrefixUV + dayStr
	utils.RedisCli.PFAdd(uvKey, ip)
}

func GetDAU(day time.Time) int64 {
	dayStr := day.Format("20060102")
	auKey := constant.RedisKeyPrefixAU + dayStr
	count, err := utils.RedisCli.BitCount(auKey, BitCountRange).Result()
	if err != nil {
		log.Printf("获取DAU出错")
		return 0
	}
	return count
}

func GetIntervalAU(start time.Time, end time.Time) int64 {
	if !start.Before(end) || end.Sub(start) > time.Hour*24*40 {
		return 0
	}
	destKey := constant.RedisKeyPrefixAU + start.Format("20060102") + "-" + end.Format("20060102")
	count, err := utils.RedisCli.BitCount(destKey, BitCountRange).Result()
	if err == nil && count > 0 {
		return count
	}
	var dayStrList []string
	for s := start; !s.After(end); s = s.Add(time.Hour*24) {
		dayStr := constant.RedisKeyPrefixAU + s.Format("20060102")
		dayStrList = append(dayStrList, dayStr)
	}
	err = utils.RedisCli.BitOpOr(destKey, dayStrList...).Err()
	if err != nil {
		return 0
	}
	count, err = utils.RedisCli.BitCount(destKey, BitCountRange).Result()
	if err != nil {
		log.Printf("获取InternalDAU出错")
		return 0
	}
	return count
}

func GetDUV(day time.Time) int64 {
	dayStr := day.Format("20060102")
	uvKey := constant.RedisKeyPrefixUV + dayStr
	count, err := utils.RedisCli.PFCount(uvKey).Result()
	if err != nil {
		log.Printf("获取DUV出错")
		return 0
	}
	return count
}

func GetIntervalUV(start time.Time, end time.Time) int64 {
	if !start.Before(end) {
		return 0
	}
	destKey := constant.RedisKeyPrefixUV + start.Format("20060102") + "-" + end.Format("20060102")
	count, err := utils.RedisCli.PFCount(destKey).Result()
	if err == nil && count > 0 {
		return count
	}
	var dayStrList []string
	for s := start; !s.After(end); s.Add(time.Hour*24) {
		dayStr := constant.RedisKeyPrefixUV + s.Format("20060102")
		dayStrList = append(dayStrList, dayStr)
	}
	err = utils.RedisCli.PFMerge(destKey, dayStrList...).Err()
	if err != nil {
		return 0
	}
	count, err = utils.RedisCli.PFCount(destKey).Result()
	if err != nil {
		log.Printf("获取InternalDUV出错")
		return 0
	}
	return count
}
