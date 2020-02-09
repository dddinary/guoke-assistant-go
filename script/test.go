package main

import (
	"guoke-helper-golang/service"
	"log"
)
func main() {
	log.Printf("#########################################")
	err := service.UpdateLectureList()
	log.Printf("%+v\n", err)
}

