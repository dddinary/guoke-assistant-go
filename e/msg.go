package e

import "github.com/gin-gonic/gin"

var errMsgs = map[int]string {
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",

	ErrorLoginFailed		: "登录失败",
	ErrorGetCaptchaFailed	: "获取验证码失败",

	CLOSING: "该功能正在维护",
}


func GetMsg(code int) string {
	msg, ok := errMsgs[code]
	if ok {
		return msg
	}
	return errMsgs[ERROR]
}

func ErrResp(code int) gin.H {
	return gin.H{
		"status": code,
		"message": GetMsg(code),
	}
}
