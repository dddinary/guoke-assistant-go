package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/service"
	"guoke-assistant-go/utils"
	"net/http"
)

func GetComingLectures(c *gin.Context) {
	lectures := service.GetComingLectures()
	if lectures == nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, lectures)
}

func GetLecture(c *gin.Context) {
	lid := utils.ValidateInt(c.DefaultQuery("lid", ""), 10)
	if lid <= 0 {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	res, err := service.GetLecture(lid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, res)
	return
}
