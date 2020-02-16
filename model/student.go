package model

import (
	"crypto/sha1"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"guoke-assistant-go/config"
	"guoke-assistant-go/constant"
	"guoke-assistant-go/utils"
	"log"
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

func FindStudentById(sid int) (*Student, error) {
	var (
		student	Student
	)
	db.Where("id = ?", sid).First(&student)
	if student.Id != sid {
		return nil, ErrorStudentNotFound
	}
	return &student, nil
}

func FindStudentsByIdList(sidList []int) ([]Student, error) {
	var (
		err			error
		students	[]Student
	)
	err = db.Where("id in (?)", sidList).Find(&students).Error
	if err != nil {
		return students, err
	}
	return students, nil
}

func FindStudentByAccount(account string) (*Student, error) {
	var (
		student	Student
	)
	db.Where("account = ?", account).First(&student)
	if student.Id == 0 {
		return nil, ErrorStudentNotFound
	}
	return &student, nil
}

func FindStudentByToken(token string) (*Student, error) {
	var (
		err			error
		student		Student
		studentPtr	*Student
	)
	studentPtr, err = GetStudentByTokenFromRedis(token)
	if err == nil && studentPtr != nil && studentPtr.Token == token {
		return studentPtr, nil
	}
	db.Where("token = ?", token).First(&student)
	if student.Token != token {
		return nil, ErrorStudentNotFound
	}
	_ = AddTokenStudentToRedis(token, &student)
	return &student, nil
}

func (student *Student) UpdateToken() string {
	trx := db.Begin()
	defer trx.Commit()

	token := genToken(student.Openid)
	_ = DeleteTokenStudentInRedis(student.Token)
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

	trx.Set("gorm:query_option", "FOR UPDATE").
		Where("account = ?", account).First(&student)
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

func GetStudentByTokenFromRedis(token string) (*Student, error) {
	var(
		err			error
		student		Student
		studentStr	string
	)
	tokenKey := constant.RedisKeyPrefixToken + token
	studentStr, err = utils.RedisCli.Get(tokenKey).Result()
	if err != nil {
		log.Printf("redis中token获取student出错：%v\n", err)
		return nil, err
	}
	err = jsoniter.UnmarshalFromString(studentStr, &student)
	if err != nil || student.Token != token {
		log.Printf("redis中token获取student反序列化出错：%v\n", err)
		return nil, err
	}
	return &student, nil
}

func AddTokenStudentToRedis(token string, student *Student) error {
	var(
		err			error
		studentStr	string
	)
	tokenKey := constant.RedisKeyPrefixToken + token
	studentStr, err = jsoniter.MarshalToString(*student)
	if err != nil {
		log.Printf("redis中存入token-student序列化出错：%v\n", err)
		return err
	}
	err = utils.RedisCli.Set(tokenKey, studentStr, 0).Err()
	if err != nil {
		log.Printf("edis中存入token-student出错：%v\n", err)
		return err
	}
	return nil
}

func DeleteTokenStudentInRedis(token string) error {
	tokenKey := constant.RedisKeyPrefixToken + token
	err := utils.RedisCli.Del(tokenKey).Err()
	if err != nil {
		log.Printf("redis中token-Student删除出错：%v\n", err)
		return err
	}
	return nil
}

func UpdateStudentBlockStatus(uid, status int) error {
	var (
		err			error
		student		Student
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()
	trx.Where("id = ?", uid).First(&student)
	if student.Id != uid || student.Status == status {
		trx.Rollback()
		return nil
	}
	if err = trx.Model(&student).Updates(map[string]interface{}{"status": status}).
		Error; err != nil {
		trx.Rollback()
		return err
	}
	// 给被禁用或者解禁的账户发通知
	if status == constant.StudentStatusBlocked {
		_ = addNotificationInTrx(trx, 0, 0, uid, constant.NotificationKindAdminBlock, "")
	} else if status == constant.StudentStatusCommon {
		_ = addNotificationInTrx(trx, 0, 0, uid, constant.NotificationKindAdminUnblock, "")
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	_ = DeleteTokenStudentInRedis(student.Token)
	return nil
}
