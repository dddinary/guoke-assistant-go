package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"guoke-assistant-go/config"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/service"
	"guoke-assistant-go/utils"
	"net/http"
	"strconv"
	"time"
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

func AdminNotification(c *gin.Context) {
	uid := utils.ValidateInt(c.DefaultQuery("uid", ""), 10)
	content := c.DefaultQuery("content", "")
	if content == "" {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	err := service.AdminAddNotification(uid, content)
	if err != nil {
		logrus.Println(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}

func AdminGetDAU(c *gin.Context) {
	dayStr := utils.ValidateInt(c.DefaultQuery("day", ""), 8)
	day, err := time.Parse("20060102", strconv.Itoa(dayStr))
	if err != nil {
		logrus.Println("Here is :", err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	dau := service.GetDAU(day)
	c.JSON(http.StatusOK, map[string]int64{"dau": dau})
	return
}

func AdminGetIntervalDAU(c *gin.Context) {
	startStr := utils.ValidateInt(c.DefaultQuery("start", ""), 8)
	endStr := utils.ValidateInt(c.DefaultQuery("end", ""), 8)
	dayStart, err := time.Parse("20060102", strconv.Itoa(startStr))
	if err != nil {
		logrus.Println(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	dayEnd, err := time.Parse("20060102", strconv.Itoa(endStr))
	if err != nil {
		logrus.Println(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	aus := service.GetIntervalAU(dayStart, dayEnd)
	c.JSON(http.StatusOK, map[string]int64{"aus": aus})
	return
}
