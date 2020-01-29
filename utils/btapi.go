package utils

import (
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strings"
)

type ImageJsonResp struct {
	Code		string
	ImgUrl		string
	Width		string
	Height		string
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func BTGetWallPaperUrl() string {
	return getImageUrl("http://api.btstu.cn/sjbz/api.php?lx=fengjing&format=json")
}

func BTGetAvatarUrl() string {
	return getImageUrl("http://api.btstu.cn/sjtx/api.php?lx=c3&format=json")
}

func getImageUrl(api string) string {
	resp, err := http.Get(api)
	if err != nil || resp.StatusCode != http.StatusOK {
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	avatarInfo := ImageJsonResp{}
	err = json.Unmarshal(body, &avatarInfo)
	if err != nil {
		return ""
	}
	url := strings.Replace(avatarInfo.ImgUrl, "large", "small", 1)
	return url
}
