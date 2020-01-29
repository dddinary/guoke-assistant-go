package service

import (
	"encoding/json"
	"fmt"
	"guoke-helper-golang/config"
	"io/ioutil"
	"log"
	"net/http"
)

type CodeSessionResp struct {
	Openid		string
	SessionKey	string
	UnionId		string
	ErrCode		int
	ErrMsg		string
}

func CodeToSession(code string) string {
	urlFormat := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	resp, err := http.Get(fmt.Sprintf(urlFormat, config.WeChatConf.AppId, config.WeChatConf.AppSecret, code))
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	sessionInfo := CodeSessionResp{}
	err = json.Unmarshal(body, &sessionInfo)
	if err != nil || sessionInfo.ErrCode != 0 {
		return ""
	}
	return sessionInfo.Openid
}
