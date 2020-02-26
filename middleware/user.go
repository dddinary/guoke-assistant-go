package middleware

import (
	"github.com/gin-gonic/gin"
	"guoke-assistant-go/config"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/model"
	"net/http"
)


func GetReqUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.DefaultQuery("_t", "")
		var uid int
		var blocked bool
		if token != "" && len(token) < 100 {
			student, _ := model.FindStudentByToken(token)
			if student != nil {
				uid = student.Id
				if student.Status == 1 {
					blocked = true
				} else {
					blocked = false
				}
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
			c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
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
			c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
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
			c.JSON(http.StatusOK, constant.ErrResp(constant.BANNED))
		} else {
			c.Next()
		}
	}
}