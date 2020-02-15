package model

import (
	"guoke-assistant-go/constant"
	"time"
)

type PostLike struct {
	Id 			int			`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Pid			int			`json:"pid" gorm:"type:int"`
	Uid			int			`json:"uid" gorm:"type:int"`
	CreatedAt	time.Time	`json:"created_at" gorm:"type:datetime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"type:datetime"`
	Deleted		int			`json:"deleted" gorm:"type:int"`
}

func AddPostLike(uid, pid int) error {
	var (
		err      error
		post     Post
		postLike PostLike
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()
	trx.Set("gorm:query_option", "FOR UPDATE").First(&post, pid)
	if post.Id != pid || post.Deleted == 1{
		trx.Rollback()
		return ErrorPostNotFound
	}
	trx.Where("pid = ? AND uid = ?", pid, uid).First(&postLike)
	if postLike.Id != 0 && postLike.Deleted == 0 {
		trx.Rollback()
		return nil
	}
	if postLike.Id != 0 {
		if err = trx.Model(&postLike).Updates(map[string]interface{}{"deleted": 0, "updated_at": time.Now()}).
			Error; err != nil {
				trx.Rollback()
				return err
		}
	} else {
		postLike = PostLike{Pid: pid, Uid:uid, CreatedAt:time.Now(), UpdatedAt:time.Now(), Deleted:0}
		if err = trx.Create(&postLike).Error; err != nil {
			trx.Rollback()
			return err
		}
	}
	if err = trx.Model(&post).Updates(Post{Like:post.Like + 1}).Error; err != nil {
		trx.Rollback()
		return err
	}
	// 给被赞的人发通知
	_ = addNotificationInTrx(trx, post.Id, uid, post.Uid, constant.NotificationKindLikePost)
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func DeletePostLike(uid, pid int) error {
	var (
		err		error
		post	Post
		like	PostLike
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()
	trx.Set("gorm:query_option", "FOR UPDATE").First(&post, pid)
	if post.Id != pid {
		trx.Rollback()
		return ErrorPostNotFound
	}
	trx.Where("pid = ? AND uid = ?", pid, uid).First(&like)
	if like.Id == 0 || like.Deleted == 1 {
		trx.Rollback()
		return nil
	}
	if err = trx.Model(&like).Updates(map[string]interface{}{"deleted": 1, "updated_at": time.Now()}).
		Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Model(&post).Updates(Post{Like:post.Like - 1}).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func IfLikedPost(uid, pid int) bool {
	var (
		err      error
		postLike PostLike
	)
	if err = db.Where("pid = ? AND uid = ?", pid, uid).Find(&postLike).Error; err != nil {
		return false
	}
	if postLike.Id > 0 && postLike.Deleted == 0 {
		return true
	}
	return false
}

func FindPostLikesByPostId(pid int) ([]PostLike, error) {
	var (
		err			error
		postLikes	[]PostLike
	)
	if err = db.Where("pid = ? AND deleted = ?", pid, 0).Find(&postLikes).Error; err != nil {
		return nil, err
	}
	return postLikes, nil
}
