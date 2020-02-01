package middleware

import (
	"github.com/gin-gonic/gin"
	"guoke-helper-golang/model"
	"log"
	"net/http"
)

const UidKey = "reqUid"

func GetReqUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		var stu *model.Student
		var uid int
		if token != "" && len(token) < 20 {
			stu, _ = model.FindStudentByToken(token)
		}
		if stu != nil {
			log.Printf("Current user: %v", *stu)
			uid = stu.Id
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
			c.JSON(http.StatusUnauthorized, "Token is not valid")
		} else {
			c.Next()
		}
	}
}