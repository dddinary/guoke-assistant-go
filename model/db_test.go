package model

import (
	"guoke-helper-golang/utils"
	"testing"
)

func TestSomething(t *testing.T) {
	posts, _ := FindPostsByCondition(-1, 0, 0, 5)
	for _, post := range posts {
		t.Log(utils.StructToMap(&post))
	}
}
