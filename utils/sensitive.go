package utils

import (
	"github.com/importcjj/sensitive"
	"guoke-assistant-go/constant"
	"os"
)

var SensFilter *sensitive.Filter

func init() {
	SensFilter = sensitive.New()
	dictFile := constant.SensitiveDictPath
	for i := 0; i < 5; i++ {
		if _, err := os.Stat(dictFile); err == nil {
			break
		} else {
			dictFile = "../" + dictFile
		}
	}
	if err := SensFilter.LoadWordDict(dictFile); err != nil {
		panic(err)
	}
}
