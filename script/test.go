package main

import (
	"fmt"
	"guoke-helper-golang/utils"
	"log"
)
func main() {
	log.Printf("#########################################")
	for i := 0; i < 10; i++ {
		fmt.Println("\"" + utils.BTGetWallPaperUrl() + "\",")
	}
}

