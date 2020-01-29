package config

import (
	"testing"
)


func TestReadConfig(t *testing.T) {
	t.Logf("%s\n", AppConf.Name)
	t.Logf("%v\n", AppConf.Release)
	t.Logf("%s\n", MysqlConf.Host)
	t.Logf("%s\n", RedisConf.Host)
	t.Logf("%s\n", WeChatConf.AppId)
	t.Logf("%s\n", LogConf.LogfilePath)
}
