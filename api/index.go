package api

import (
	"github.com/gin-gonic/gin"
	"guoke-helper-golang/constant"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, constant.ErrResp(constant.SUCCESS))
	return
}
