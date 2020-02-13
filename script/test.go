package main

import (
	"guoke-helper-golang/model"
	"log"
)
func main() {
	log.Printf("#########################################")
	res := model.IfLikedPost(2, 1)
	log.Println(res)
}

