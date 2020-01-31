package model

import (
	"errors"
	"time"
)

type Post struct {
	Id 			int32		`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Uid			int32		`json:"uid" gorm:"type:int"`
	Content		string		`json:"content" gorm:"type:text"`
	Kind		int32		`json:"kind" gorm:"type:int"`
	Like		int32		`json:"like" gorm:"type:int"`
	View		int32		`json:"view" gorm:"type:int"`
	Comment		int32		`json:"comment" gorm:"type:int"`
	CreatedAt	time.Time	`json:"created_at" gorm:"type:datetime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"type:datetime"`
	Deleted		int32		`json:"deleted" gorm:"type:int"`
}

var ErrorPostNotFound = errors.New("没找到对应的Post")

func AddPost(uid int32, content string, kind int32) error {
	var (
		err		error
		post	Post
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	post = Post{Uid:uid, Content:content, Kind:kind, Like:0, View:0, Comment:0,
		CreatedAt:time.Now(), UpdatedAt:time.Now(), Deleted:0}
	if err = trx.Create(&post).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func DeletePost(uid, pid int32) error {
	var (
		err		error
		post	Post
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	if err = trx.Set("gorm:query_option", "FOR UPDATE").First(&post, pid).Error; err != nil {
		trx.Rollback()
		return err
	}
	if post.Id != pid {
		trx.Rollback()
		return ErrorPostNotFound
	}
	if err = trx.Model(&post).Updates(Post{Deleted:1}).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func FindPostById(pid int32) (*Post, error) {
	var (
		err		error
		post	Post
	)
	if err = db.First(&post, pid).Error; err != nil {
		return nil, err
	}
	if post.Id != pid {
		return nil, ErrorPostNotFound
	}
	return &post, nil
}

func FindPostsByIdList(idList []int32) ([]Post, error) {
	var (
		err		error
		posts	[]Post
	)
	if err = db.Where("id in (?)", idList).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}