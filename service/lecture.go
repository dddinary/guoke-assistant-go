package service

import (
	"guoke-helper-golang/model"
	"time"
)

func GetLecture() map[string][]model.Lecture {
	return model.GetComingLectures()
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