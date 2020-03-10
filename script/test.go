package main

import (
	"guoke-assistant-go/utils"
	"log"
	"time"
)
func main() {
	log.Printf("#########################################")
	utils.BotMsgNewUser(0, "测试用户", "家里蹲大学")
	time.Sleep(time.Second * 3)
}




