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

func CommentPost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	content := c.DefaultQuery("content", "")
	if content == "" || len(content) > config.AppConf.CommentMaxLen {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
	}
	uid := c.MustGet(UidKey).(int)
	err := service.CommentPost(uid, pid, content)
	if err != nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}

func CommentComment(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	cid := utils.ValidateInt(c.DefaultQuery("cid", ""), 10)
	content := c.DefaultQuery("content", "")
	if content == "" || len(content) > config.AppConf.CommentMaxLen {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
	}
	uid := c.MustGet(UidKey).(int)
	err := service.CommentComment(uid, pid, cid, content)
	if err != nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}

func LikeComment(c *gin.Context) {
	cid := utils.ValidateInt(c.DefaultQuery("cid", ""), 10)
	uid := c.MustGet(UidKey).(int)
	err := service.LikeComment(uid, cid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}

func UnlikeComment(c *gin.Context) {
	cid := utils.ValidateInt(c.DefaultQuery("cid", ""), 10)
	uid := c.MustGet(UidKey).(int)
	err := service.UnlikeComment(uid, cid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}

func DeleteComment(c *gin.Context) {
	cid := utils.ValidateInt(c.DefaultQuery("cid", ""), 10)
	uid := c.MustGet(UidKey).(int)
	err := service.DeleteComment(uid, cid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}