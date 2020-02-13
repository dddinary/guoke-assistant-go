package api

import (
	"github.com/gin-gonic/gin"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/service"
	"net/http"
)

func WxLogin(c *gin.Context) {
	code := c.DefaultPostForm("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	openid := service.CodeToSession(code)
	if openid == "" {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorGetCaptchaFailed))
		return
	}
	c.JSON(http.StatusOK, gin.H{"openid": openid})
}
