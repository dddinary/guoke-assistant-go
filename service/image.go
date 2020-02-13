package service

import "guoke-helper-golang/model"

func GetImagesByPostId(pid int) ([]string, error) {
	return model.FindImagesByPostId(pid)
}
