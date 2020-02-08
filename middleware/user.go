package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"guoke-helper-golang/config"
	"guoke-helper-golang/e"
	"guoke-helper-golang/model"
	"guoke-helper-golang/utils"
	"net/http"
)

const UidKey = "reqUid"

func GetReqUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		var stu *model.Student
		var uid int
		if token != "" && len(token) < 100 {
			uid = utils.ValidateTokenByRedis(token)
			if uid == 0 {
				stu, _ = model.FindStudentByToken(token)
				if stu != nil {
					logrus.Printf("Current user: %v", *stu)
					uid = stu.Id
					_ = utils.AddTokenToRedis(token, uid)
				}
			}
		}
		c.Set(UidKey, uid)
		c.Next()
	}
}

func NeedLogin() gin.HandlerFunc {

	return func(c *gin.Context) {
		uid := c.MustGet(UidKey).(int)
		if uid == 0 {
			c.Abort()
			c.JSON(http.StatusUnauthorized, e.ErrResp(e.ErrorInvalidParams))
		} else {
			c.Next()
		}
	}
}

func AdminOnly() gin.HandlerFunc {

	return func(c *gin.Context) {
		uid := c.MustGet(UidKey).(int)
		if uid != config.AppConf.Admin {
			c.Abort()
			c.JSON(http.StatusUnauthorized, e.ErrResp(e.ErrorInvalidParams))
		} else {
			c.Next()
		}
	}
}