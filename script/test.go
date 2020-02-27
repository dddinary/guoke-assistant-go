package main

import (
	"guoke-assistant-go/utils"
	"log"
)
func main() {
	log.Printf("#########################################")
	res, _ := utils.GetCosCredential()
	log.Println(res)
}

