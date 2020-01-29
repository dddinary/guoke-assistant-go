package api

import (
	"github.com/gin-gonic/gin"
	"guoke-helper-golang/config"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": config.AppConf.Name,
	})
}
