package model

import (
	"errors"
	"time"
)

type Comment struct {
	Id 			int32		`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Pid			int32		`json:"pid" gorm:"type:int"`
	Uid			int32		`json:"uid" gorm:"type:int"`
	Content		string		`json:"content" gorm:"type:text"`
	Kind		int32		`json:"kind" gorm:"type:int"`
	Like		int32		`json:"like" gorm:"type:int"`
	CreatedAt	time.Time	`json:"created_at" gorm:"type:datetime"`
	Deleted		int32		`json:"deleted" gorm:"type:int"`
}

func AddComment(pid, uid int32, content string, kind, like int32) error {
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

	if err = trx.Set("gorm:query_option", "FOR UPDATE").First(&post, pid).Error; err != nil {
		trx.Rollback()
		return err
	}
	if post.Id == 0 || post.Deleted == 1 {
		trx.Rollback()
		return ErrorPostNotFound
	}
	comment := Comment{Pid:pid, Uid:uid, Content:content, Kind:kind, Like:like, CreatedAt:time.Now(), Deleted:0}
	if err = trx.Create(&comment).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Model(&post).Update("comment", post.Comment + 1).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func DeleteComment(commentId, uid int32) error {
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
	if comment.Uid != uid {
		trx.Rollback()
		return errors.New("不是对应的用户")
	}
	if err = trx.First(&post, comment.Pid).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Model(&post).Update(Post{Comment:post.Comment - 1}).Error; err != nil {
		trx.Rollback()
		return err
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

func FindCommentsByPostId(pid int32) ([]Comment, error) {
	var (
		err			error
		comments	[]Comment
	)
	if err = db.Where("pid = ? AND deleted = ?", pid, 0).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
