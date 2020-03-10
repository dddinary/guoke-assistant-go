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

func BotMsgNewUser(uid int, name, dpt string) {
	AsyncMsgToBot("新用户", fmt.Sprintf("uid:%d 姓名：%s 院系：%s", uid, name, dpt))
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
	data := map[string]string{
		"title": title,
		"text": text,
	}
	header := req.Header{
		"Content-Type": "application/json",
	}
	botHook := config.BotConf.Data
	_, err := req.Post(botHook, req.BodyJSON(data), header)
	if err != nil {
		log.Println(err)
	}
}
