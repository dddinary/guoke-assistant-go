package router

import (
	"github.com/gin-gonic/gin"
	"guoke-assistant-go/api"
	"guoke-assistant-go/middleware"
)

func InitRouterEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.LoggerToFile(), middleware.GetReqUser())

	r.GET("/", api.Index)
	// r.GET("/getCaptcha", api.GetCaptcha)
	// r.POST("/wxLogin", api.WxLogin)
	r.POST("/loginCourse", api.LoginAndGetCourse)

	r.GET("/getLectures", api.GetComingLectures)
	r.GET("/getLecture", api.GetLecture)
	r.GET("/getNews", api.GetNews)
	r.GET("/getPost", api.GetPost)
	r.GET("/searchPost", api.SearchPost)
	r.GET("/getUserPost", api.GetUserPost)
	r.GET("/getPostLikes", api.GetPostLikes)
	r.GET("/getPostComments", api.GetPostComments)
	r.GET("/getPostImages", api.GetPostImages)
	r.GET("/getStudentInfo", api.GetStudentInfo)
	r.GET("/getStudentsInfoList", api.GetStudentsInfoList)

	needLogin := r.Group("/s", middleware.NeedLogin())

	needLogin.GET("/tempCredential", middleware.Blocker(), api.GetCosCredential)
	needLogin.GET("/publish", middleware.Blocker(), api.Publish)
	needLogin.GET("/commentPost", middleware.Blocker(), api.CommentPost)
	needLogin.GET("/commentComment", middleware.Blocker(), api.CommentComment)

	needLogin.GET("/changeAvatar", api.ChangeAvatar)
	needLogin.GET("/getStarPost", api.GetStaredPost)
	needLogin.GET("/likePost", api.LikePost)
	needLogin.GET("/unlikePost", api.UnlikePost)
	needLogin.GET("/likeComment", api.LikeComment)
	needLogin.GET("/unlikeComment", api.UnlikeComment)
	needLogin.GET("/starPost", api.StarPost)
	needLogin.GET("/unstarPost", api.UnstarPost)
	needLogin.GET("/deletePost", api.DeletePost)
	needLogin.GET("/deleteComment", api.DeleteComment)
	needLogin.GET("/messageCount", api.CountUnreadNotifications)
	needLogin.GET("/getMessage", api.GetOnesNotifications)
	needLogin.GET("/readMessage", api.MarkReadNotifications)
	needLogin.GET("/deleteMessage", api.DeleteNotifications)

	adminOnly := r.Group("/a", middleware.AdminOnly())

	adminOnly.GET("/deletePost", api.AdminDeletePost)
	adminOnly.GET("/deleteComment", api.AdminDeleteComment)
	adminOnly.GET("/blockStudent", api.AdminBlockStudent)
	adminOnly.GET("/unblockStudent", api.AdminUnblockStudent)
	adminOnly.GET("/notify", api.AdminNotification)
	return r
}