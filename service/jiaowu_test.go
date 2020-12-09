package service

import "testing"

func TestLoginAndGetCourse(t *testing.T) {
	t.Log(LoginAndGetCourse("", "dddinary@163.com", "tobeno.1", ""))
}

func TestUpdateLectureList(t *testing.T) {
	t.Log(UpdateLectureList())
}
