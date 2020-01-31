package model

import "time"

type CommentLike struct {
	Id 			int32		`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Cid			int32		`json:"cid" gorm:"type:int"`
	Uid			int32		`json:"uid" gorm:"type:int"`
	CreatedAt	time.Time	`json:"created_at" gorm:"type:datetime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"type:datetime"`
	Deleted		int32		`json:"deleted" gorm:"type:int"`
}

func AddCommentLike(cid, uid int32) error {
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
	if err = trx.Set("gorm:query_option", "FOR UPDATE").First(&comment, cid).Error; err != nil {
		trx.Rollback()
		return err
	}
	if comment.Id != cid || comment.Deleted == 1{
		trx.Rollback()
		return ErrorCommentNotFound
	}
	if err = trx.Where("cid = ? AND uid = ?", cid, uid).First(&commentLike).Error; err != nil {
		trx.Rollback()
		return err
	}
	if commentLike.Id != 0 && commentLike.Deleted == 0 {
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
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func DeleteCommentLike(pid, uid int32) error {
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
	if err = trx.Set("gorm:query_option", "FOR UPDATE").First(&comment, pid).Error; err != nil {
		trx.Rollback()
		return err
	}
	if comment.Id != pid {
		trx.Rollback()
		return ErrorPostNotFound
	}
	if err = trx.Where("pid = ? AND uid = ?", pid, uid).First(&commentLike).Error; err != nil {
		trx.Rollback()
		return err
	}
	if commentLike.Id == 0 || commentLike.Deleted == 1 {
		trx.Rollback()
		return nil
	}
	if err = trx.Model(&commentLike).Updates(map[string]interface{}{"deleted": 1, "updated_at": time.Now()}).
		Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Model(&comment).Updates(Post{Like: comment.Like - 1}).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}
