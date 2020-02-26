package api

import (
	"github.com/gin-gonic/gin"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/service"
	"net/http"
)

func GetCaptcha(c *gin.Context) {
	openid := c.DefaultQuery("openid", "")
	if openid == "" {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	imgData := service.GetCaptcha(openid)
	if imgData == nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorGetCaptchaFailed))
		return
	}
	c.Data(http.StatusOK, "image/jpeg", imgData)
	return
}

func LoginAndGetCourse(c *gin.Context) {
	openid := c.DefaultPostForm("openid", "")
	username := c.DefaultPostForm("username", "")
	pwd := c.DefaultPostForm("pwd", "")
	avatar := c.DefaultPostForm("avatar", "")
	if username == "" || pwd == "" {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	res := service.LoginAndGetCourse(openid, username, pwd, avatar)
	if res == nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorLoginFailed))
		return
	}
	c.JSON(http.StatusOK, res)
	return
}
