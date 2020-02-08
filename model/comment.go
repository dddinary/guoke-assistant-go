package model

import (
	"errors"
	"guoke-helper-golang/config"
	"time"
)

type Comment struct {
	Id 			int			`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Pid			int			`json:"pid" gorm:"type:int"`
	Cid			int			`json:"cid" gorm:"type:int"`
	Uid			int			`json:"uid" gorm:"type:int"`
	Content		string		`json:"content" gorm:"type:text"`
	Like		int			`json:"like" gorm:"type:int"`
	CreatedAt	time.Time	`json:"created_at" gorm:"type:datetime"`
	Deleted		int			`json:"deleted" gorm:"type:int"`
}

var ErrorCommentNotFound = errors.New("没找到相应的Comment")

// 如果cid不为0就是二级评论，否则就是一级评论
func AddComment(uid, pid, cid int, content string) error {
	var (
		err error
		post Post
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	if cid == 0 {
		trx = trx.Set("gorm:query_option", "FOR UPDATE")
	}
	if err = trx.First(&post, pid).Error; err != nil {
		trx.Rollback()
		return err
	}
	if post.Id == 0 || post.Deleted == 1 {
		trx.Rollback()
		return ErrorPostNotFound
	}
	comment := Comment{Pid:pid, Uid:uid, Cid:cid, Content:content, Like:0, CreatedAt:time.Now(), Deleted:0}
	if err = trx.Create(&comment).Error; err != nil {
		trx.Rollback()
		return err
	}
	if cid == 0 {
		if err = trx.Model(&post).Update("comment", post.Comment + 1).Error; err != nil {
			trx.Rollback()
			return err
		}
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func DeleteComment(uid, commentId int) error {
	var (
		err		error
		comment	Comment
		post	Post
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	if err = trx.Set("gorm:query_option", "FOR UPDATE").First(&comment, commentId).Error; err != nil {
		trx.Rollback()
		return err
	}
	if comment.Uid != uid && uid != config.AppConf.Admin {
		trx.Rollback()
		return errors.New("不是对应的用户")
	}
	if err = trx.First(&post, comment.Pid).Error; err != nil {
		trx.Rollback()
		return err
	}
	if comment.Cid == 0 {
		if err = trx.Model(&post).Update(Post{Comment:post.Comment - 1}).Error; err != nil {
			trx.Rollback()
			return err
		}
	}
	if err = trx.Model(&comment).Updates(Comment{Deleted:1}).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func FindCommentsByPostId(pid int) ([]Comment, error) {
	var (
		err			error
		comments	[]Comment
	)
	if err = db.Where("pid = ? AND deleted = ?", pid, 0).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
