package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"guoke-assistant-go/config"
	"guoke-assistant-go/job"
	"guoke-assistant-go/router"
	"log"
)

func main() {
	if config.AppConf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	c := cron.New()
	id, err := c.AddJob("@hourly", job.LectureJob{})
	if err != nil {
		log.Fatalf("lecture job 创建失败: %+v\n", err)
	}
	log.Printf("lecture job 创建成功 id=%+v", id)
	c.Start()

	r := router.InitRouterEngine()
	err = r.Run(":" + config.AppConf.Port)
	if err != nil {
		log.Fatalf("主服务启动失败: %+v\n", err)
	}
}

