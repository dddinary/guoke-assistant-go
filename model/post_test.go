package model

import (
	"guoke-assistant-go/utils"
	"log"
	"testing"
)

func TestAddPostToES(t *testing.T) {
	for i := 1; i <= 90; i++ {
		post, _ := FindPostById(i)
		log.Printf("%v\n", utils.AddPostToES(post.Id, post.Uid, post.Content, post.CreatedAt, post.Deleted))
	}
}

func TestSearchPost(t *testing.T) {
	pidList, _ := utils.SearchPostInES("程序课", 0, 20)
	posts, _ := FindPostsByIdList(pidList)
	for _, post := range posts {
		log.Printf("%v\n", post)
	}
}
