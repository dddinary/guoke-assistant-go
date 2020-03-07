package main

import (
	"guoke-assistant-go/utils"
	"log"
)
func main() {
	log.Printf("#########################################")
	res := utils.SensFilter.Replace("日本av帝国", '*')
	log.Println(res)
}

