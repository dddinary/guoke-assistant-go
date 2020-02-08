package api

import (
	"github.com/gin-gonic/gin"
	"guoke-helper-golang/e"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, e.ErrResp(e.SUCCESS))
}
