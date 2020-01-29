package api

import (
	"github.com/gin-gonic/gin"
	"guoke-helper-golang/service"
	"net/http"
)

func GetLecture(c *gin.Context) {
	lectures := service.GetLecture()
	if lectures == nil || len(lectures) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数非法",
		})
		return
	}
	c.JSON(http.StatusOK, lectures)
}
