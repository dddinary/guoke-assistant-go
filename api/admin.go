package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"guoke-assistant-go/config"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/service"
	"guoke-assistant-go/utils"
	"net/http"
)

func AdminDeletePost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	err := service.DeletePost(uid, pid)
	if err != nil {
		logrus.Println(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}

func AdminDeleteComment(c *gin.Context) {
	cid := utils.ValidateInt(c.DefaultQuery("cid", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	err := service.DeleteComment(uid, cid)
	if err != nil {
		logrus.Println(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}

func AdminBlockStudent(c *gin.Context) {
	uid := utils.ValidateInt(c.DefaultQuery("uid", ""), 10)
	if uid == config.AdminConf.Uid {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	err := service.BlockStudent(uid)
	if err != nil {
		logrus.Println(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}

func AdminUnblockStudent(c *gin.Context) {
	uid := utils.ValidateInt(c.DefaultQuery("uid", ""), 10)
	err := service.UnblockStudent(uid)
	if err != nil {
		logrus.Println(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}
