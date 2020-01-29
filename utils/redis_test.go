package utils

import "testing"

func TestRedis(t *testing.T) {
	t.Log(RedisCli.Get("name").Val())
}
