package model

import "testing"

func TestFindStudentByAccount(t *testing.T) {
	t.Log(FindStudentByAccount("dddinary@163.com"))
}
