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
		c.JSON(http.StatusBadRequest, e.ErrResp(e.InvalidParams))
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
	cap := c.DefaultQuery("cap", "")
	if openid == "" || username == "" || pwd == "" || cap == "" {
		c.JSON(http.StatusBadRequest, e.ErrResp(e.InvalidParams))
		return
	}
	res := service.LoginAndGetCourse(openid, username, pwd, cap)
	if res == nil {
		c.JSON(http.StatusInternalServerError, e.ErrResp(e.ErrorLoginFailed))
	}
	c.JSON(http.StatusOK, res)
}
