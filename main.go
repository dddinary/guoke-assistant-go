package main

import (
	"github.com/gin-gonic/gin"
	"guoke-helper-golang/api"
	"guoke-helper-golang/config"
	"guoke-helper-golang/middleware"
)

func main() {
	// config.InitConfig("config/develop.yaml")
	// logrus.Info("载入配置成功")

	r := gin.Default()
	r.Use(middleware.GetReqUser())

	r.GET("/", api.Index)
	// r.GET("/getCaptcha", api.GetCaptcha)
	r.GET("/loginCourse", api.LoginAndGetCourse)
	r.GET("/getLecture", api.GetLecture)
	r.GET("/wxLogin", api.WxLogin)

	r.GET("/getNews", api.GetNews)
	r.GET("/getPost", api.GetPost)
	r.GET("/getUserPost", api.GetUserPost)

	needLogin := r.Group("/s", middleware.NeedLogin())

	needLogin.GET("/getStarPost", api.GetStaredPost)
	needLogin.GET("/publish", api.Publish)
	needLogin.GET("/commentPost", api.CommentPost)
	needLogin.GET("/commentComment", api.CommentComment)
	needLogin.GET("/likePost", api.LikePost)
	needLogin.GET("/unlikePost", api.UnlikePost)
	needLogin.GET("/likeComment", api.LikeComment)
	needLogin.GET("/unlikeComment", api.UnlikeComment)
	needLogin.GET("/starPost", api.StarPost)
	needLogin.GET("/unstarPost", api.UnstarPost)
	needLogin.GET("/deletePost", api.DeletePost)
	needLogin.GET("/deleteComment", api.DeleteComment)

	adminOnly := r.Group("/a", middleware.AdminOnly())
	adminOnly.GET("/deletePost", api.AdminDeletePost)
	adminOnly.GET("/deleteComment", api.AdminDeleteComment)

	r.Run(":" + config.AppConf.Port)
}

