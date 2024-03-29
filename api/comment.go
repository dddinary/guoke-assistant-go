package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/service"
	"guoke-assistant-go/utils"
	"net/http"
)

func CommentPost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	content := c.DefaultQuery("content", "")
	if content == "" || len(content) > constant.CommentMaxLen || pid <= 0 {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	content = utils.SensFilter.Replace(content, '*')
	uid := c.MustGet(constant.ContextKeyUid).(int)
	err := service.CommentPost(uid, pid, content)
	if err != nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}

func CommentComment(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	cid := utils.ValidateInt(c.DefaultQuery("cid", ""), 10)
	ruid := utils.ValidateInt(c.DefaultQuery("ruid", ""), 10)
	content := c.DefaultQuery("content", "")
	if content == "" || len(content) > constant.CommentMaxLen || pid <= 0 || cid <= 0 {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	content = utils.SensFilter.Replace(content, '*')
	uid := c.MustGet(constant.ContextKeyUid).(int)
	err := service.CommentComment(uid, pid, cid, ruid, content)
	if err != nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}

func LikeComment(c *gin.Context) {
	cid := utils.ValidateInt(c.DefaultQuery("cid", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	if cid <= 0 {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	err := service.LikeComment(uid, cid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}

func UnlikeComment(c *gin.Context) {
	cid := utils.ValidateInt(c.DefaultQuery("cid", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	if cid <= 0 {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	err := service.UnlikeComment(uid, cid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}

func DeleteComment(c *gin.Context) {
	cid := utils.ValidateInt(c.DefaultQuery("cid", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	if cid <= 0 {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	err := service.DeleteComment(uid, cid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}