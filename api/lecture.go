package api

import (
	"github.com/gin-gonic/gin"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/service"
	"net/http"
)

func GetLecture(c *gin.Context) {
	lectures := service.GetLecture()
	if lectures == nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ERROR))
		return
	}
	c.JSON(http.StatusOK, lectures)
}
