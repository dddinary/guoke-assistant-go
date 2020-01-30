package model

import "testing"

func TestSomething(t *testing.T) {
	student := FindStudentByAccount("dddinary@163.com")
	t.Log(student)
}
