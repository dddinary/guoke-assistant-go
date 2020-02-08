package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"guoke-helper-golang/config"
	"time"
)

type Post struct {
	Id 			int			`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Uid			int			`json:"uid" gorm:"type:int"`
	Content		string		`json:"content" gorm:"type:text"`
	Kind		int			`json:"kind" gorm:"type:int"`
	Like		int			`json:"like" gorm:"type:int"`
	View		int			`json:"view" gorm:"type:int"`
	Comment		int			`json:"comment" gorm:"type:int"`
	CreatedAt	time.Time	`json:"created_at" gorm:"type:datetime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"type:datetime"`
	Deleted		int			`json:"deleted" gorm:"type:int"`
}

var ErrorPostNotFound = errors.New("没找到对应的Post")

func (post *Post) ToMap() map[string]interface{} {
	res := make(map[string]interface{})
	res["id"] = post.Id

	return res
}

func AddPost(uid int, content string, kind int) error {
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

func DeletePost(uid, pid int) error {
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
	if post.Uid != uid && uid != config.AppConf.Admin {
		trx.Rollback()
		return ErrorNotRightUser
	}
	if err = trx.Model(&post).Updates(Post{Deleted:1, UpdatedAt:time.Now()}).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func FindPostById(pid int) (*Post, error) {
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

func FindPostsByIdList(idList []int) ([]Post, error) {
	var (
		err		error
		posts	[]Post
	)
	if err = db.Where("id in (?)", idList).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

var orderMap = map[int]string{
	0: "created_at desc",
	1: "like desc",
	2: "comment desc",
}

func FindPostsByCondition(kind, order, pageIdx, pageSize int) ([]Post, error) {
	var (
		err		error
		posts	[]Post
		handler *gorm.DB
	)
	if kind > 0 {
		handler = db.Where("kind = ?", kind)
	} else {
		handler = db
	}
	orderStr, ok := orderMap[order]
	if !ok {
		orderStr = orderMap[0]
	}
	if err = handler.Order(orderStr).Offset(pageIdx*pageSize).Limit(pageSize).Find(&posts).Error; err != nil {
			return nil, err
	}
	return posts, nil
}

func FindPostsByUid(uid, pageIdx, pageSize int) ([]Post, error) {
	var (
		err		error
		posts	[]Post
	)
	if err = db.Where("uid = ?", uid).Order("created_at desc").
		Offset(pageIdx*pageSize).Limit(pageSize).Find(&posts).Error; err != nil {
			return nil, err
	}
	return posts, nil
}

func FindStaredPosts(uid, pageIdx, pageSize int) ([]Post, error) {
	var (
		err		error
		posts	[]Post
		stars	[]Star
		idList	[]int
	)
	if err = db.Where("uid = ?", uid).Order("created_at desc").
		Offset(pageIdx*pageSize).Limit(pageSize).Find(&stars).Error; err != nil {
			return nil, err
	}
	idList = []int{}
	for _, rec := range stars {
		idList = append(idList, rec.Pid)
	}
	posts, err = FindPostsByIdList(idList)
	return posts, err
}