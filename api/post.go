package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"guoke-helper-golang/config"
	"guoke-helper-golang/constant"
	"guoke-helper-golang/service"
	"guoke-helper-golang/utils"
	"net/http"
)

const UidKey = "reqUid"

func GetNews(c *gin.Context) {
	kind := utils.ValidateInt(c.DefaultQuery("kind", ""), 10)
	order := utils.ValidateInt(c.DefaultQuery("order", ""), 10)
	pageIdx := utils.ValidateInt(c.DefaultQuery("page", ""), 10)

	uid := c.MustGet(UidKey).(int)

	res, err := service.GetNews(uid, kind, order, pageIdx)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, res)
}

func GetPost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(UidKey).(int)
	res, err := service.GetPostDetail(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, res)
}

func GetUserPost(c *gin.Context) {
	wantedUid := utils.ValidateInt(c.DefaultQuery("uid", ""), 10)
	pageIdx := utils.ValidateInt(c.DefaultQuery("page", ""), 10)
	uid := c.MustGet(UidKey).(int)
	res, err := service.GetUserPost(uid, wantedUid, pageIdx)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, res)
}

func GetStaredPost(c *gin.Context) {
	pageIdx := utils.ValidateInt(c.DefaultQuery("page", ""), 10)
	uid := c.MustGet(UidKey).(int)
	res, err := service.GetStaredPost(uid, pageIdx)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, res)
}

func Publish(c *gin.Context) {
	kind := utils.ValidateInt(c.DefaultQuery("kind", ""), 10)
	content := c.DefaultQuery("content", "")
	if content == "" || len(content) > config.AppConf.PostMaxLen {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
	}
	uid := c.MustGet(UidKey).(int)
	err := service.AddPost(uid, content, kind)
	if err != nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}

func LikePost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(UidKey).(int)
	err := service.LikePost(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}

func UnlikePost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(UidKey).(int)
	err := service.UnlikePost(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}

func StarPost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(UidKey).(int)
	err := service.StarPost(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}

func UnstarPost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(UidKey).(int)
	err := service.UnstarPost(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}

func DeletePost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(UidKey).(int)
	err := service.DeletePost(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}