package model

import (
	"errors"
	"time"
)

type Lecture struct {
	Id 			int32		`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Lid			int32		`json:"lid" gorm:"type:int"`
	Name		string		`json:"name" gorm:"type:varchar(255)"`
	Category	int32		`json:"category" gorm:"type:int"`
	Dpt			string		`json:"dpt" gorm:"type:varchar(255)"`
	Start		time.Time	`json:"start" gorm:"type:datetime"`
	End			time.Time	`json:"end" gorm:"type:datetime"`
	Venue		string		`json:"venue" gorm:"type:varchar(255)"`
	Desc		string		`json:"desc" gorm:"type:text"`
	Pic			string		`json:"pic" gorm:"type:varchar(255)"`
}

var ErrorLectureHasExist = errors.New("该lecture已存在")

func GetComingLectures() map[string][]Lecture {
	var (
		hum, sci	[]Lecture
	)
	lectures := make(map[string][]Lecture)
	loc, _ := time.LoadLocation("Local")
	from := time.Date(2019, 11, 12, 0, 0, 0, 0, loc)
	db.Where("start >= ? and category = ?", from, 2).Find(&hum)
	db.Where("start >= ? and category = ?", from, 1).Find(&sci)
	lectures["humanity"] = hum
	lectures["science"] = sci
	return lectures
}

func AddLecture(lid int32, name string, category int32, dpt string, start, end time.Time, venue, desc, pic string) error {
	var (
		err		error
		lecture	Lecture
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	if err = trx.First(&lecture, lid).Error; err != nil {
		trx.Rollback()
		return err
	}
	if lecture.Id == lid {
		trx.Rollback()
		return ErrorLectureHasExist
	}
	lecture = Lecture{Lid:lid, Name:name, Category:category, Dpt:dpt,
		Start:start, End:end, Venue:venue, Desc:desc, Pic:pic}
	if err = trx.Create(&lecture).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}