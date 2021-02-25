package utils

import (
	"fmt"
	"github.com/imroc/req"
	"guoke-assistant-go/config"
	"log"
)

func BotMsgWarning(content string) {
	AsyncMsgToBot("警报！！！！！！！！！", content)
}

func BotMsgUserLogin(uid int, name, dpt string) {
	AsyncMsgToBot("用户登录", fmt.Sprintf("uid:%d 姓名：%s 院系：%s", uid, name, dpt))
}

func BotMsgJobStart(jobName string) {
	AsyncMsgToBot("定时任务启动", jobName)
}

func BotMsgLectureUpdate(lid int, name string) {
	AsyncMsgToBot("新添加讲座", fmt.Sprintf("lid:%d 名称：%s", lid, name))
}

func AsyncMsgToBot(title, text string) {
	go MsgToBot(title, text)
}

func MsgToBot(title, text string) {
	var err error
	header := req.Header{
		"Content-Type": "application/json",
	}
	data := map[string]string{
		"title": title,
		"text": text,
	}

	feishuBotHook := config.BotConf.Feishu
	_, err = req.Post(feishuBotHook, req.BodyJSON(data), header)
	if err != nil {
		log.Println(err)
	}

	ddData := map[string]interface{} {
		"msgtype": "text",
		"text": map[string]string {
			"content": "【果壳助手】\n" + title + ":" + text + "",
		},
	}
	dingdingBotHook := config.BotConf.Dingding
	_, err = req.Post(dingdingBotHook, req.BodyJSON(ddData), header)
	if err != nil {
		log.Println(err)
	}
}
