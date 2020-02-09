package model

import (
	"errors"
)

type Course struct {
	Cid			int		`json:"cid" gorm:"primary_key"`
	Code		string	`json:"code" gorm:"type:varchar(255)"`
	Dpt			string	`json:"dpt" gorm:"type:varchar(255)"`
	Name		string	`json:"name" gorm:"type:varchar(255)"`
	Attr		string	`json:"attr" gorm:"type:varchar(255)"`
	Major		string	`json:"major" gorm:"type:varchar(255)"`
	Keshi		string	`json:"keshi" gorm:"type:varchar(255)"`
	Score		string	`json:"score" gorm:"type:varchar(255)"`
	Shouke		string	`json:"shouke" gorm:"type:varchar(255)"`
	Kaoshi		string	`json:"kaoshi" gorm:"type:varchar(255)"`
	Professor	string	`json:"professor" gorm:"type:varchar(255)"`
	Lecturer	string	`json:"lecturer" gorm:"type:varchar(255)"`
	Assistant	string	`json:"assistant" gorm:"type:varchar(255)"`
}

var ErrorCourseNotFound = errors.New("没有找到相应的课程")

func FindCourseByCid(cid int) (*Course, error) {
	var (
		course	Course
	)
	db.Where("cid = ?", cid).First(&course)
	if course.Cid == 0 {
		return nil, ErrorCourseNotFound
	}
	return &course, nil
}

func FindCoursesByCidList(cidList []int) ([]Course, error) {
	var (
		err		error
		courses	[]Course
	)
	if err = db.Where("cid in (?)", cidList).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}