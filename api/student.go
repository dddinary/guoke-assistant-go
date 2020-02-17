package api

import (
	"github.com/gin-gonic/gin"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/service"
	"guoke-assistant-go/utils"
	"net/http"
	"strings"
)

func GetStudentInfo(c *gin.Context) {
	uid := utils.ValidateInt(c.DefaultQuery("sid", ""), 10)
	stuInfo, err := service.GetStudentNoSecretInfoById(uid)
	if err != nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	c.JSON(http.StatusOK, stuInfo)
}

func GetStudentsInfoList(c *gin.Context) {
	sidListStr:= c.DefaultQuery("sids", "")
	var sidStrList []string
	if sidListStr == "" || len(sidListStr) > 100 {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	} else {
		sidStrList = strings.Split(sidListStr, ",")
	}
	var sidList []int
	for _, sidStr := range sidStrList {
		sid := utils.ValidateInt(sidStr, 10)
		if sid > 0 {
			sidList = append(sidList, sid)
		}
	}
	stuInfoMap, err := service.GetStudentsNoSecretInfoByIdList(sidList)
	if err != nil {
		c.JSON(http.StatusOK, constant.ErrResp(constant.ErrorInvalidParams))
		return
	}
	c.JSON(http.StatusOK, stuInfoMap)
}
