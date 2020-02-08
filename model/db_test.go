package model

import (
	"testing"
)

func TestSomething(t *testing.T) {
	stu, err := FindStudentByToken("98c876002effd5ecd796d4a06da5766706a44c71")
	if err != nil {
		t.Log(err)
	}
	t.Log(stu)
}
