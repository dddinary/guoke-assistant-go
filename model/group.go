package model

import (
	"crypto/md5"
	"fmt"
	"log"
)

type Group struct {
	Id 			uint		`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Account		string		`json:"account" gorm:"type:varchar(255)"`
	Password	string		`json:"password" gorm:"type:varchar(255)"`
	Name		string		`json:"name" gorm:"type:varchar(255)"`
	Avatar		string		`json:"avatar" gorm:"type:varchar(255)"`
	Status		int32		`json:"status" gorm:"type:int"`
}

func FindGroupById(cid int32) *Group {
	trx := db.Begin()
	defer trx.Commit()

	group := new(Group)
	trx.Where("id = ?", cid).First(group)
	if group.Id == 0 {
		return nil
	}
	return group
}

func FindGroupByAccount(account string) *Group {
	trx := db.Begin()
	defer trx.Commit()

	group := new(Group)
	trx.Where("account = ?", account).First(group)
	if group.Id == 0 {
		return nil
	}
	return group
}

func AddGroup(account, password, name, avatar string) bool {
	trx := db.Begin()
	defer trx.Commit()

	group := Group{}
	trx.Set("gorm:query_option", "FOR UPDATE").
		Where("account = ?", account).First(&group)
	if group.Id != 0 {
		return false
	}

	hash := md5.New()
	hash.Write([]byte(password))
	hashedPwd := fmt.Sprintf("%x", hash.Sum(nil))
	group = Group{Account:account, Password:hashedPwd, Name:name, Avatar:avatar}
	if err := trx.Create(&group).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func CheckGroupPwd(account, password string) bool {
	trx := db.Begin()
	defer trx.Commit()

	hash := md5.New()
	hash.Write([]byte(password))
	hashedPwd := fmt.Sprintf("%x", hash.Sum(nil))
	group := Group{}
	trx.Where("account = ?", account).First(&group)
	if group.Id != 0 && group.Password == hashedPwd {
		return true
	}
	return false
}