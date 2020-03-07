package service

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req"
	"guoke-assistant-go/config"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/utils"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

type CheckSecResp struct {
	Errcode	int
	ErrMsg	string
}

func CheckMsgSecurity(content string) bool {
	errLogMsg := "检测内容安全"
	data := map[string]string{"content":content}
	urlFormat := "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s"
	resp, err := req.Post(fmt.Sprintf(urlFormat, GetAccessToken()), req.BodyJSON(data))
	if err != nil {
		log.Printf("%s：%+v\n", errLogMsg, err)
		return false
	}
	secResp := CheckSecResp{}
	if err = resp.ToJSON(&secResp); err != nil {
		log.Printf("%s：%+v\n", errLogMsg, err)
		return false
	}
	return secResp.Errcode != 87014
}
// 这个方法有问题！！！
//func CheckImageSecurity(imgUrl string) bool {
//	errLogMsg := "检测内容安全"
//	imgResp, err := req.Get(imgUrl)
//	if err != nil {
//		return false
//	}
//	data := imgResp.Bytes()
//	urlFormat := "https://api.weixin.qq.com/wxa/img_sec_check?access_token=%s"
//	resp, err := req.Post(fmt.Sprintf(urlFormat, GetAccessToken()), data)
//	if err != nil {
//		log.Printf("%s：%+v\n", errLogMsg, err)
//		return false
//	}
//	secResp := CheckSecResp{}
//	if err = resp.ToJSON(&secResp); err != nil {
//		log.Printf("%s：%+v\n", errLogMsg, err)
//		return false
//	}
//	log.Println(secResp)
//	return secResp.Errcode != 87014
//}


func GetAccessToken() string {
	acToken := getAccessTokenFromRedis()
	if acToken == "" {
		acToken = getAccessTokenFromWeChat()
	}
	return acToken
}

type AccessTokenResp struct {
	AccessToken	string	`json:"access_token"`
	ExpiresIn	int		`json:"expires_in"`
	ErrCode		int
	ErrMsg		string
}

func getAccessTokenFromWeChat() string {
	urlFormat := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	resp, err := http.Get(fmt.Sprintf(urlFormat, config.WeChatConf.AppId, config.WeChatConf.AppSecret))
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	acTokenInfo := AccessTokenResp{}
	err = json.Unmarshal(body, &acTokenInfo)
	if err != nil || acTokenInfo.ErrCode != 0 {
		return ""
	}
	err = utils.RedisCli.Set(constant.RedisKeyAccessToken, acTokenInfo.AccessToken, time.Hour).Err()
	return acTokenInfo.AccessToken
}

func getAccessTokenFromRedis() string {
	acToken, err := utils.RedisCli.Get(constant.RedisKeyAccessToken).Result()
	if err != nil {
		return ""
	}
	return acToken
}