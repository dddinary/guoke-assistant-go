package main

import (
	"github.com/gin-gonic/gin"
	"guoke-helper-golang/api"
	"guoke-helper-golang/config"
)

func main() {
	// config.InitConfig("config/develop.yaml")
	// logrus.Info("载入配置成功")

	r := gin.Default()
	r.GET("/", api.Index)
	r.GET("/getCaptcha", api.GetCaptcha)
	r.GET("/loginCourse", api.LoginAndGetCourse)
	r.GET("/getLecture", api.GetLecture)
	r.GET("/wxLogin", api.WxLogin)

	r.Run(":" + config.AppConf.Port)
}

