package model

import (
	"crypto/md5"
	"errors"
	"fmt"
)

type Group struct {
	Id 			int			`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Account		string		`json:"account" gorm:"type:varchar(255)"`
	Password	string		`json:"password" gorm:"type:varchar(255)"`
	Name		string		`json:"name" gorm:"type:varchar(255)"`
	Avatar		string		`json:"avatar" gorm:"type:varchar(255)"`
	Status		int			`json:"status" gorm:"type:int"`
}

var ErrorGroupNotFound = errors.New("没有找到相应的group")
var ErrorGroupHasExist = errors.New("已经存在该group")

func FindGroupById(gid int) (*Group, error) {
	var (
		group	Group
	)
	db.First(&group, gid)
	if group.Id == 0 {
		return nil, ErrorGroupNotFound
	}
	return &group, nil
}

func FindGroupByAccount(account string) (*Group, error) {
	var (
		group	Group
	)
	db.Where("account = ?", account).First(&group)
	if group.Id == 0 {
		return nil, ErrorGroupNotFound
	}
	return &group, nil
}

func AddGroup(account, password, name, avatar string) error {
	var (
		err		error
		group	Group
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	trx.Set("gorm:query_option", "FOR UPDATE").
		Where("account = ?", account).First(&group)
	if group.Id != 0 {
		trx.Rollback()
		return ErrorGroupHasExist
	}

	hash := md5.New()
	hash.Write([]byte(password))
	hashedPwd := fmt.Sprintf("%x", hash.Sum(nil))
	group = Group{Account:account, Password:hashedPwd, Name:name, Avatar:avatar}
	if err = trx.Create(&group).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func CheckGroupPwd(account, password string) bool {
	var (
		group	Group
	)
	hash := md5.New()
	hash.Write([]byte(password))
	hashedPwd := fmt.Sprintf("%x", hash.Sum(nil))
	if err := db.Where("account = ?", account).First(&group).Error; err != nil {
		return false
	}
	if group.Id != 0 && group.Password == hashedPwd {
		return true
	}
	return false
}