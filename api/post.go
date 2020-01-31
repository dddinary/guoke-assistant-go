package api

import (
	"github.com/gin-gonic/gin"
	"guoke-helper-golang/model"
	"guoke-helper-golang/service"
	"guoke-helper-golang/utils"
	"net/http"
)

func GetNews(c *gin.Context) {
	token := c.DefaultQuery("token", "")
	kind := utils.ValidateInt(c.DefaultQuery("kind", ""), 10)
	order := utils.ValidateInt(c.DefaultQuery("order", ""), 10)
	pageIdx := utils.ValidateInt(c.DefaultQuery("page", ""), 10)

	var stu *model.Student
	var uid int

	if token != "" && len(token) < 20 {
		stu, _ = model.FindStudentByToken(token)
	}
	if stu != nil {
		uid = stu.Id
	}
	res, _ := service.GetNews(uid, kind, order, pageIdx)
	c.JSON(http.StatusOK, res)
}
