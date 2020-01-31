package model

import "testing"

func TestSomething(t *testing.T) {
	student, _ := FindStudentByAccount("dddinary@163.com")
	t.Log(student)
	posts, _ := FindPostsByIdList([]int32{1})
	t.Log(posts)
}
