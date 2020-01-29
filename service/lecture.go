package service

import "guoke-helper-golang/model"

func GetLecture() map[string][]model.Lecture {
	return model.GetComingLectures()
}
