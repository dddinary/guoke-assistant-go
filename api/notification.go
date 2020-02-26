package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/service"
	"guoke-assistant-go/utils"
	"net/http"
	"strings"
)

func CountUnreadNotifications(c *gin.Context) {
	uid := c.MustGet(constant.ContextKeyUid).(int)
	count := service.GetUnreadNotificationsCount(uid)
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func GetOnesNotifications(c *gin.Context) {
	pageIdx := utils.ValidateInt(c.DefaultQuery("page", ""), 10)
	uid := c.MustGet(constant.ContextKeyUid).(int)
	res, err := service.GetOnesNotifications(uid, pageIdx)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, res)
	return
}

func MarkReadNotifications(c *gin.Context) {
	nidStr := c.DefaultQuery("nid", "")
	var nidStrList []string
	if nidStr != "" && len(nidStr) < 100 {
		nidStrList = strings.Split(nidStr, ",")
	}
	var nidList []int
	for _, nidStr := range nidStrList {
		nid := utils.ValidateInt(nidStr, 10)
		if nid > 0 {
			nidList = append(nidList, nid)
		}
	}
	uid := c.MustGet(constant.ContextKeyUid).(int)
	err := service.MarkReadNotifications(uid, nidList)
	if err != nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}

func DeleteNotifications(c *gin.Context) {
	nidStr := c.DefaultQuery("nid", "")
	var nidStrList []string
	if nidStr != "" && len(nidStr) < 100 {
		nidStrList = strings.Split(nidStr, ",")
	}
	var nidList []int
	for _, nidStr := range nidStrList {
		nid := utils.ValidateInt(nidStr, 10)
		if nid > 0 {
			nidList = append(nidList, nid)
		}
	}
	uid := c.MustGet(constant.ContextKeyUid).(int)
	err := service.DeleteNotifications(uid, nidList)
	if err != nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
}
