package model

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"guoke-helper-golang/config"
	"log"
	"time"
)

type Student struct {
	Id 			uint		`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Account		string		`json:"account" gorm:"type:varchar(255)"`
	Name		string		`json:"name" gorm:"type:varchar(255)"`
	Dpt			string		`json:"dpt" gorm:"type:varchar(255)"`
	Avatar		string		`json:"avatar" gorm:"type:varchar(255)"`
	Openid		string		`json:"openid" gorm:"type:varchar(255)"`
	Token		string		`json:"token" gorm:"type:varchar(255)"`
	Status		int32		`json:"status" gorm:"type:int"`
}

func FindStudentById(cid int32) *Student {
	trx := db.Begin()
	defer trx.Commit()

	student := new(Student)
	trx.Where("id = ?", cid).First(student)
	if student.Id == 0 {
		return nil
	}
	return student
}

func FindStudentByAccount(account string) *Student {
	trx := db.Begin()
	defer trx.Commit()

	student := new(Student)
	trx.Where("account = ?", account).First(student)
	if student.Id == 0 {
		return nil
	}
	return student
}

func FindStudentByToken(token string) *Student {
	trx := db.Begin()
	defer trx.Commit()

	student := new(Student)
	trx.Where("token = ?", token).First(student)
	if student.Id == 0 {
		return nil
	}
	return student
}

func (student *Student) UpdateToken() string {
	trx := db.Begin()
	defer trx.Commit()

	token := genToken(student.Openid)
	trx.Model(&student).Update("token", token)
	return token
}

func AddStudent(account, name, dpt, avatar, openid string) (string, error) {
	trx := db.Begin()
	defer trx.Commit()

	student := Student{}
	trx.Set("gorm:query_option", "FOR UPDATE").
		Where("account = ?", account).First(&student)
	if student.Id != 0 {
		return "", errors.New("该用户已存在")
	}

	token := genToken(openid)
	student = Student{Account:account, Name:name, Dpt:dpt, Avatar:avatar, Openid:openid, Token:token}
	if err := trx.Create(&student).Error; err != nil {
		log.Println(err)
		return "", err
	}
	return token, nil
}

func genToken(openid string) string {
	s := time.Now().String() + openid + config.AppConf.Magic
	h := sha1.New()
	h.Write([]byte(s))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}