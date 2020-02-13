package main

import (
	"guoke-assistant-go/model"
	"log"
)
func main() {
	log.Printf("#########################################")
	res := model.IfLikedPost(2, 1)
	log.Println(res)
}

