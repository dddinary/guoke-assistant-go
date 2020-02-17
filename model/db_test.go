package model

import (
	"testing"
)

func TestSomething(t *testing.T) {
	var (
		star	Star
	)
	db.Where("pid = ? AND uid = ?", 10, 2).Find(&star)
	t.Logf("%+v\n", star)
}
