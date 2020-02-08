package api

import (
	"github.com/gin-gonic/gin"
	"guoke-helper-golang/e"
	"guoke-helper-golang/service"
	"net/http"
)

func GetCaptcha(c *gin.Context) {
	openid := c.DefaultQuery("openid", "")
	if openid == "" {
		c.JSON(http.StatusBadRequest, e.ErrResp(e.ErrorInvalidParams))
		return
	}
	imgData := service.GetCaptcha(openid)
	if imgData == nil {
		c.JSON(http.StatusBadRequest, e.ErrResp(e.ErrorGetCaptchaFailed))
		return
	}
	c.Data(http.StatusOK, "image/jpeg", imgData)
}

func LoginAndGetCourse(c *gin.Context) {
	openid := c.DefaultQuery("openid", "")
	username := c.DefaultQuery("username", "")
	pwd := c.DefaultQuery("pwd", "")
	avatar := c.DefaultQuery("avatar", "")
	if openid == "" || username == "" || pwd == "" {
		c.JSON(http.StatusBadRequest, e.ErrResp(e.ErrorInvalidParams))
		return
	}
	res := service.LoginAndGetCourse(openid, username, pwd, avatar)
	if res == nil {
		c.JSON(http.StatusInternalServerError, e.ErrResp(e.ErrorLoginFailed))
	}
	c.JSON(http.StatusOK, res)
}
