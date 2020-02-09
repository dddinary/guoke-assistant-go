package middleware

import (
	"github.com/gin-gonic/gin"
	"guoke-helper-golang/config"
	"guoke-helper-golang/constant"
	"guoke-helper-golang/model"
	"net/http"
)


func GetReqUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		var stu *model.Student
		var uid int
		var blocked bool
		if token != "" && len(token) < 100 {
			stu, _ = model.FindStudentByToken(token)
			uid = stu.Id
			if stu.Status == 1 {
				blocked = true
			} else {
				blocked = false
			}
		}
		c.Set(constant.ContextKeyUid, uid)
		c.Set(constant.ContextKeyBlocked, blocked)
		c.Next()
	}
}

func NeedLogin() gin.HandlerFunc {

	return func(c *gin.Context) {
		uid := c.MustGet(constant.ContextKeyUid).(int)
		if uid == 0 {
			c.Abort()
			c.JSON(http.StatusUnauthorized, constant.ErrResp(constant.ErrorInvalidParams))
		} else {
			c.Next()
		}
	}
}

func AdminOnly() gin.HandlerFunc {

	return func(c *gin.Context) {
		uid := c.MustGet(constant.ContextKeyUid).(int)
		if uid != config.AdminConf.Uid {
			c.Abort()
			c.JSON(http.StatusUnauthorized, constant.ErrResp(constant.ErrorInvalidParams))
		} else {
			c.Next()
		}
	}
}

func Blocker() gin.HandlerFunc {

	return func(c *gin.Context) {
		blocked := c.MustGet(constant.ContextKeyBlocked).(bool)
		if blocked {
			c.Abort()
			c.JSON(http.StatusUnauthorized, constant.ErrResp(constant.BANNED))
		} else {
			c.Next()
		}
	}
}