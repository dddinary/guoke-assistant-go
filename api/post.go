package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/service"
	"guoke-assistant-go/utils"
	"net/http"
)

func GetNews(c *gin.Context) {
	kind := utils.ValidateInt(c.DefaultQuery("kind", ""), 10)
	order := utils.ValidateInt(c.DefaultQuery("order", ""), 10)
	pageIdx := utils.ValidateInt(c.DefaultQuery("page", ""), 10)

	uid := c.MustGet(constant.ContextKeyUid).(int)

	res, err := service.GetNews(uid, kind, order, pageIdx)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func GetPost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	res, err := service.GetPostDetail(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func GetPostComments(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	res, err := service.GetCommentsByPostId(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func GetPostLikes(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	res, err := service.GetLikesByPostId(pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func GetPostImages(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	res, err := service.GetImagesByPostId(pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func GetUserPost(c *gin.Context) {
	wantedUid := utils.ValidateInt(c.DefaultQuery("uid", ""), 10)
	pageIdx := utils.ValidateInt(c.DefaultQuery("page", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	res, err := service.GetUserPost(uid, wantedUid, pageIdx)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func GetStaredPost(c *gin.Context) {
	pageIdx := utils.ValidateInt(c.DefaultQuery("page", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	res, err := service.GetStaredPost(uid, pageIdx)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func Publish(c *gin.Context) {
	kind := utils.ValidateInt(c.DefaultQuery("kind", ""), 10)
	content := c.DefaultQuery("content", "")
	images, ok := c.GetQueryArray("images")
	if !ok {
		images = []string{}
	}
	if content == "" && len(images) > 0 {
		content = "发表图片"
	}
	if content == "" || len(content) > constant.PostMaxLen {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	uid := c.MustGet(constant.ContextKeyUid).(int)
	err := service.AddPost(uid, content, kind, images)
	if err != nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}

func LikePost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	err := service.LikePost(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}

func UnlikePost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	err := service.UnlikePost(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}

func StarPost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	err := service.StarPost(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}

func UnstarPost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	err := service.UnstarPost(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}

func DeletePost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	err := service.DeletePost(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}