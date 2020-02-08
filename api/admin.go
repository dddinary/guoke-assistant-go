package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"guoke-helper-golang/e"
	"guoke-helper-golang/service"
	"guoke-helper-golang/utils"
	"net/http"
)

func AdminDeletePost(c *gin.Context) {
	pid := utils.ValidateInt(c.DefaultQuery("pid", ""), 10)
	uid := c.MustGet(UidKey).(int)
	err := service.DeletePost(uid, pid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, e.ErrResp(e.ERROR))
	}
	c.JSON(http.StatusOK, e.ErrResp(e.SUCCESS))
}

func AdminDeleteComment(c *gin.Context) {
	cid := utils.ValidateInt(c.DefaultQuery("cid", ""), 10)
	uid := c.MustGet(UidKey).(int)
	err := service.DeleteComment(uid, cid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, e.ErrResp(e.ERROR))
	}
	c.JSON(http.StatusOK, e.ErrResp(e.SUCCESS))
}
