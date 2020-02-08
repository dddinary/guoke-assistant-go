package api

import (
	"github.com/gin-gonic/gin"
	"guoke-helper-golang/e"
	"guoke-helper-golang/service"
	"net/http"
)

func WxLogin(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, e.ErrResp(e.ErrorInvalidParams))
		return
	}
	openid := service.CodeToSession(code)
	if openid == "" {
		c.JSON(http.StatusOK, e.ErrResp(e.ErrorGetCaptchaFailed))
		return
	}
	c.JSON(http.StatusOK, gin.H{"openid": openid})
}
