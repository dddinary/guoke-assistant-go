package utils

import (
	"log"
	"testing"
	"time"
)

func TestESSearch(t *testing.T) {
	pidList, err := SearchPostInES("程序课", 0)
	log.Printf("%v\n%v\n", pidList, err)
	log.Printf("%v\n", AddPostToES(100, 100, "200", time.Now(), 0))
	log.Panicln(MarkPostInESDeleted(100))
}
