package constant

import "github.com/gin-gonic/gin"

const (
	SUCCESS       = 200
	ERROR         = 500

	ErrorInvalidParams = 400
	ErrorUnauthorized = 401

	ErrorLoginFailed		= 10001
	ErrorGetCaptchaFailed	= 10002

	CLOSING = 50001
	BANNED = 50002
	Limited = 5003
)

var errMsgs = map[int]string {
	SUCCESS:            "ok",
	ERROR:              "fail",

	ErrorInvalidParams: "请求参数错误",
	ErrorUnauthorized: "用户未认证或认证错误",

	ErrorLoginFailed		: "登录失败",
	ErrorGetCaptchaFailed	: "获取验证码失败",

	CLOSING: "该功能正在维护",
	BANNED: "该用户已被限制",
	Limited: "该用户已被限流",
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
