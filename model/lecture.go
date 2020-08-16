package model

import (
	"errors"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/utils"
	"log"
	"time"
)

type Lecture struct {
	Id 			int			`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Lid			int			`json:"lid" gorm:"type:int"`
	Name		string		`json:"name" gorm:"type:varchar(255)"`
	Category	int			`json:"category" gorm:"type:int"`
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
		err 		error
		allLecture	[]Lecture
		hum, sci	[]Lecture
		lectures	map[string][]Lecture
	)
	lectures, err = GetLecturesFromRedis()
	if err == nil && lectures != nil {
		return lectures
	}
	lectures = make(map[string][]Lecture)
	loc, _ := time.LoadLocation("Local")
	from := time.Date(2020, 6, 1, 0, 0, 0, 0, loc)
	// form := time.Now().AddDate(0, 0, -1)
	db.Where("start >= ?", from).Order("start desc").Find(&allLecture)
	for _, lec := range allLecture {
		if lec.Category == constant.LectureKindHumanity {
			hum = append(hum, lec)
		} else if lec.Category == constant.LectureKindScience {
			sci = append(sci, lec)
		}
	}
	lectures["humanity"] = hum
	lectures["science"] = sci
	_ = AddLecturesToRedis(lectures)
	return lectures
}

func AddLecture(lid int, name string, category int, dpt string, start, end time.Time, venue, desc, pic string) error {
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

	trx.Where("lid = ?", lid).First(&lecture)
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

func LectureExists(lid int) (bool, error) {
	var (
		err		error
		lecture	Lecture
	)
	err = db.Where("lid = ?", lid).First(&lecture).Error
	if err != nil {
		return false, err
	}
	if lecture.Id > 0 {
		return true, nil
	}
	return false, nil
}

func GetLecturesFromRedis() (map[string][]Lecture, error) {
	var (
		err				error
		lecturesStr		string
		lectures = make(map[string][]Lecture)
	)

	lecturesStr, err = utils.RedisCli.Get(constant.RedisKeyLecture).Result()
	if err == redis.Nil {
		return nil, err
	} else if err != nil {
		log.Printf("redis中获取讲座信息出错：%v\n", err)
		return nil, err
	}
	err = jsoniter.UnmarshalFromString(lecturesStr, &lectures)
	if err != nil {
		log.Printf("讲座信息反序列化出错：%v\n", err)
		return nil, err
	}
	return lectures, nil
}

func AddLecturesToRedis(lectures map[string][]Lecture) error {
	var (
		err				error
		lecturesStr		string
	)
	lecturesStr, err = jsoniter.MarshalToString(lectures)
	if err != nil {
		log.Printf("讲座信息序列化出错：%v\n", err)
		return err
	}
	err = utils.RedisCli.Set(constant.RedisKeyLecture, lecturesStr, time.Hour * 72).Err()
	if err != nil {
		log.Printf("redis中写入讲座信息出错：%v\n", err)
		return err
	}
	return nil
}

func DeleteLecturesInRedis() error {
	err := utils.RedisCli.Del(constant.RedisKeyLecture).Err()
	if err != nil {
		log.Printf("redis中删除讲座信息出错：%v\n", err)
		return err
	}
	return nil
}