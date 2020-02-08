package main

import (
	"github.com/imroc/req"
	"guoke-helper-golang/utils"
	"log"
)
func main() {
	name := utils.RedisCli.Get("test").String()
	log.Printf(name)
}

type MainLoginRes struct {
	F	bool	`json:"f"`
	Msg	string	`json:"msg"`
}

func MainLoginWithoutCaptcha(cli *req.Req, username, password string) bool {
	var (
		err			error
		resp		*req.Resp
		loginRes	MainLoginRes
		errLogMsg	= "登录失败"
	)
	_, _ = cli.Get("http://onestop.ucas.edu.cn")
	data := req.Param{
		"username": username,
		"password": password,
	}
	headers := req.Header{
		"X-Requested-With": "XMLHttpRequest",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
	}
	loginUrl := "http://onestop.ucas.edu.cn/Ajax/Login/0"
	resp, err = cli.Post(loginUrl, data, headers)
	if err != nil {
		log.Printf("%s：%+v\n", errLogMsg, err)
	}
	if err = resp.ToJSON(&loginRes); err != nil {
		log.Printf("%s：%+v\n", errLogMsg, err)
	}
	if !loginRes.F {
		log.Printf("%s：%s", errLogMsg, loginRes.Msg)
	}
	resp, err = cli.Get(loginRes.Msg)
	if err != nil {
		log.Printf("%s：%+v\n", errLogMsg, err)
	}
	return true
}
