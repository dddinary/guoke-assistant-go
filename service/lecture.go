package service

import (
	"guoke-assistant-go/model"
	"time"
)

func GetComingLectures() map[string][]model.Lecture {
	return model.GetComingLectures()
}

func GetLecture(lid int) (model.Lecture, error) {
	return model.GetLecture(lid)
}

func AddLecture(lid int, name string, category int, dpt string, start, end time.Time, venue, desc, pic string) error {
	return model.AddLecture(lid, name, category, dpt, start, end, venue, desc, pic)
}

func LectureExists(lid int) (bool, error) {
	return model.LectureExists(lid)
}

func DeleteLectureDataInRedis() error {
	return model.DeleteLecturesInRedis()
}