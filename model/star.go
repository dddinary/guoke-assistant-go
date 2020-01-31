package model

import "time"

type Star struct {
	Id 			int			`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Pid			int			`json:"pid" gorm:"type:int"`
	Uid			int			`json:"uid" gorm:"type:int"`
	CreatedAt	time.Time	`json:"created_at" gorm:"type:datetime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"type:datetime"`
	Deleted		int			`json:"deleted" gorm:"type:int"`
}

func AddStar(uid, pid int) error {
	var (
		err		error
		post	Post
		star	Star
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()
	if err = trx.First(&post, pid).Error; err != nil {
		trx.Rollback()
		return err
	}
	if post.Id != pid || post.Deleted == 1{
		trx.Rollback()
		return ErrorPostNotFound
	}
	if err = trx.Where("pid = ? AND uid = ?", pid, uid).First(&star).Error; err != nil {
		trx.Rollback()
		return err
	}
	if star.Id != 0 && star.Deleted == 0 {
		trx.Rollback()
		return nil
	}
	if star.Id != 0 {
		if err = trx.Model(&star).Updates(map[string]interface{}{"deleted": 0, "updated_at": time.Now()}).
			Error; err != nil {
			trx.Rollback()
			return err
		}
	} else {
		star = Star{Pid:pid, Uid:uid, CreatedAt:time.Now(), UpdatedAt:time.Now(), Deleted:0}
		if err = trx.Create(&star).Error; err != nil {
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

func DeleteStar(uid, pid int) error {
	var (
		err		error
		star	Star
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()
	if err = trx.Where("pid = ? AND uid = ?", pid, uid).First(&star).Error; err != nil {
		trx.Rollback()
		return err
	}
	if star.Id == 0 || star.Deleted == 1 {
		trx.Rollback()
		return nil
	}
	if err = trx.Model(&star).Updates(map[string]interface{}{"deleted": 1, "updated_at": time.Now()}).
		Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func IfStared(uid, pid int) bool {
	var (
		err		error
		star	Star
	)
	if err = db.Where("pid = AND uid = ", pid, uid, &star).Error; err != nil {
		return false
	}
	if star.Id > 0 && star.Deleted == 0 {
		return true
	}
	return false
}

