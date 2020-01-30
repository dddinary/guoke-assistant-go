package service

import "testing"

func TestLoginAndGetCourse(t *testing.T) {
	t.Log(GetCourseDetailAndTimeTable([]int32{164707, 174042}))
}
