package model

import "testing"

func TestFindCourseByCid(t *testing.T) {
	course := FindCourseByCid(12)
	if course == nil {
		t.Log("No such course!")
	} else {
		t.Log(course.ToMap())
	}
}
