package router

import (
	"github.com/gin-gonic/gin"
	"guoke-helper-golang/api"
	"guoke-helper-golang/middleware"
)

func InitRouterEngine() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.GetReqUser())

	r.GET("/", api.Index)
	// r.GET("/getCaptcha", api.GetCaptcha)
	r.POST("/wxLogin", api.WxLogin)
	r.POST("/loginCourse", api.LoginAndGetCourse)

	r.GET("/getLecture", api.GetLecture)
	r.GET("/getNews", api.GetNews)
	r.GET("/getPost", api.GetPost)
	r.GET("/getUserPost", api.GetUserPost)
	r.GET("/getPostLikes", api.GetPostLikes)
	r.GET("/getPostComments", api.GetPostComments)
	r.GET("/getPostImages", api.GetPostImages)
	r.GET("/getStudentInfo", api.GetStudentInfo)
	r.GET("/getStudentsInfoList", api.GetStudentsInfoList)

	needLogin := r.Group("/s", middleware.NeedLogin())

	needLogin.GET("/publish", api.Publish, middleware.Blocker())
	needLogin.GET("/commentPost", api.CommentPost, middleware.Blocker())
	needLogin.GET("/commentComment", api.CommentComment, middleware.Blocker())

	needLogin.GET("/getStarPost", api.GetStaredPost)
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
	adminOnly.GET("/blockStudent", api.AdminBlockStudent)
	adminOnly.GET("/unblockStudent", api.AdminUnblockStudent)
	return r
}