package api

import (
	"github.com/gin-gonic/gin"
	"guoke-helper-golang/constant"
	"guoke-helper-golang/service"
	"guoke-helper-golang/utils"
	"net/http"
)

func GetStudentInfo(c *gin.Context) {
	uid := utils.ValidateInt(c.DefaultQuery("sid", ""), 10)
	stuInfo, err := service.GetStudentById(uid)
	if err != nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	c.JSON(http.StatusOK, stuInfo)
}

func GetStudentsInfoList(c *gin.Context) {
	sidStrList, ok := c.GetQueryArray("sid")
	if !ok {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	var sidList []int
	for _, sidStr := range sidStrList {
		sid := utils.ValidateInt(sidStr, 10)
		if sid > 0 {
			sidList = append(sidList, sid)
		}
	}
	stuInfoMap, err := service.GetStudentsByIdList(sidList)
	if err != nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	c.JSON(http.StatusOK, stuInfoMap)
}
