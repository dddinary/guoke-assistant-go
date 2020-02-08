package model

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"guoke-helper-golang/config"
	"guoke-helper-golang/utils"
	"time"
)

type Student struct {
	Id 			int			`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Account		string		`json:"account" gorm:"type:varchar(255)"`
	Name		string		`json:"name" gorm:"type:varchar(255)"`
	Dpt			string		`json:"dpt" gorm:"type:varchar(255)"`
	Avatar		string		`json:"avatar" gorm:"type:varchar(255)"`
	Openid		string		`json:"openid" gorm:"type:varchar(255)"`
	Token		string		`json:"token" gorm:"type:varchar(255)"`
	Status		int			`json:"status" gorm:"type:int"`
}

var ErrorStudentNotFound = errors.New("没有找到对应用户")
var ErrorStudentHasExist = errors.New("该账号已存在")
var ErrorNotRightUser 	 = errors.New("用户非法操作")

func FindStudentById(cid int) (*Student, error) {
	var (
		err		error
		student	Student
	)
	if err = db.Where("id = ?", cid).First(&student).Error; err != nil {
		return nil, err
	}
	if student.Id != cid {
		return nil, ErrorStudentNotFound
	}
	return &student, nil
}

func FindStudentByAccount(account string) (*Student, error) {
	var (
		err		error
		student	Student
	)
	if err = db.Where("account = ?", account).First(&student).Error; err != nil {
		return nil, err
	}
	if student.Id == 0 {
		return nil, ErrorStudentNotFound
	}
	return &student, nil
}

func FindStudentByToken(token string) (*Student, error) {
	var (
		err		error
		student	Student
	)
	if err = db.Where("token = ?", token).First(&student).Error; err != nil {
		return nil, err
	}
	if student.Id == 0 {
		return nil, ErrorStudentNotFound
	}
	return &student, nil
}

func (student *Student) UpdateToken() string {
	trx := db.Begin()
	defer trx.Commit()

	token := genToken(student.Openid)
	_ = utils.DeleteTokenInRedis(student.Token)
	trx.Model(&student).Update("token", token)
	return token
}

func AddStudent(account, name, dpt, avatar, openid string) (string, error) {
	var (
		err		error
		student	Student
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	if err = trx.Set("gorm:query_option", "FOR UPDATE").
		Where("account = ?", account).First(&student).Error; err != nil {
			trx.Rollback()
			return "", err
	}
	if student.Id != 0 {
		trx.Rollback()
		return "", ErrorStudentHasExist
	}

	token := genToken(openid)
	student = Student{Account:account, Name:name, Dpt:dpt, Avatar:avatar, Openid:openid, Token:token}
	if err = trx.Create(&student).Error; err != nil {
		trx.Rollback()
		return "", err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
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