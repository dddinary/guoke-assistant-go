package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"guoke-helper-golang/config"
	"log"
)

var db *gorm.DB

func init() {
	dbType	:= "mysql"
	format	:= "%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local"
	dbHost	:= config.MysqlConf.Host
	dbPort	:= config.MysqlConf.Port
	dbUser	:= config.MysqlConf.Username
	dbPwd	:= config.MysqlConf.Password
	dbName	:= config.MysqlConf.Database
	dbChar	:= config.MysqlConf.Charset
	connStr := fmt.Sprintf(format, dbUser, dbPwd, dbHost, dbPort, dbName, dbChar)

	var err error
	db, err = gorm.Open(dbType, connStr)
	if err != nil {
		log.Fatalln("Fail to connect database!")
	}
}

func CloseDB() {
	db.Close()
}
