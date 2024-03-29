package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/service"
	"guoke-assistant-go/utils"
	"net/http"
)

func WxLogin(c *gin.Context) {
	code := c.DefaultPostForm("code", "")
	if code == "" {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	openid := service.CodeToSession(code)
	if openid == "" {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorGetCaptchaFailed))
		return
	}
	c.JSON(http.StatusOK, gin.H{"openid": openid})
}

func GetCosCredential(c *gin.Context) {
	uid := c.MustGet(constant.ContextKeyUid).(int)
	res, err := utils.GetCosCredential(uid)
	if err != nil {
		logrus.Print(err)
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, res)
	return
}
