package model

import (
	"guoke-assistant-go/constant"
	"time"
)

type CommentLike struct {
	Id 			int			`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Cid			int			`json:"cid" gorm:"type:int"`
	Uid			int			`json:"uid" gorm:"type:int"`
	CreatedAt	time.Time	`json:"created_at" gorm:"type:datetime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"type:datetime"`
	Deleted		int			`json:"deleted" gorm:"type:int"`
}

func AddCommentLike(uid, cid int) error {
	var (
		err         error
		comment     Comment
		commentLike CommentLike
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()
	trx.Set("gorm:query_option", "FOR UPDATE").First(&comment, cid)
	if comment.Id != cid || comment.Deleted == 1{
		trx.Rollback()
		return ErrorCommentNotFound
	}
	trx.Where("cid = ? AND uid = ?", cid, uid).First(&commentLike)
	// 已经点过赞或者二级评论，都无法点赞
	if commentLike.Id != 0 && commentLike.Deleted == 0 || comment.Cid != 0{
		trx.Rollback()
		return nil
	}
	if commentLike.Id != 0 {
		if err = trx.Model(&commentLike).Updates(map[string]interface{}{"deleted": 0, "updated_at": time.Now()}).
			Error; err != nil {
			trx.Rollback()
			return err
		}
	} else {
		commentLike = CommentLike{Cid: cid, Uid:uid, CreatedAt:time.Now(), UpdatedAt:time.Now(), Deleted:0}
		if err = trx.Create(&commentLike).Error; err != nil {
			trx.Rollback()
			return err
		}
	}
	if err = trx.Model(&comment).Updates(Post{Like:comment.Like + 1}).Error; err != nil {
		trx.Rollback()
		return err
	}
	// 给被赞的人发通知
	_ = addNotificationInTrx(trx, comment.Pid, uid, comment.Uid, constant.NotificationKindLikeComment, comment.Content)
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func DeleteCommentLike(uid, cid int) error {
	var (
		err         error
		comment     Comment
		commentLike CommentLike
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()
	trx.Set("gorm:query_option", "FOR UPDATE").First(&comment, cid)
	if comment.Id != cid {
		trx.Rollback()
		return ErrorCommentNotFound
	}
	trx.Where("cid = ? AND uid = ?", cid, uid).First(&commentLike)
	if commentLike.Id == 0 || commentLike.Deleted == 1 {
		trx.Rollback()
		return nil
	}
	if err = trx.Model(&commentLike).Updates(map[string]interface{}{"deleted": 1, "updated_at": time.Now()}).
		Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Model(&comment).Updates(Comment{Like: comment.Like - 1}).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func IfLikedComment(uid, cid int) bool {
	var (
		err         error
		commentLike CommentLike
	)
	if err = db.Where("cid = AND uid = ", cid, uid, &commentLike).Error; err != nil {
		return false
	}
	if commentLike.Id > 0 && commentLike.Deleted == 0 {
		return true
	}
	return false
}