package service

import "guoke-assistant-go/model"

func GetImagesByPostId(pid int) ([]string, error) {
	return model.FindImagesByPostId(pid)
}
