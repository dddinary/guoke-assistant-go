package model

import (
	"guoke-assistant-go/constant"
	"log"
	"time"
)

type Notification struct {
	Id 			int			`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Pid			int			`json:"pid" gorm:"type:int"`
	Notifier	int			`json:"notifier" gorm:"type:int"`
	Receiver	int			`json:"receiver" gorm:"type:int"`
	Kind		int			`json:"kind" gorm:"type:int"`
	Status		int			`json:"status" gorm:"type:int"`
	CreatedAt	time.Time	`json:"created_at" gorm:"type:datetime"`
}

func AddNotification(pid, notifier, receiver, kind int) error {
	var (
		err				error
		notification	Notification
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()
	notification = Notification{Pid:pid, Notifier:notifier, Receiver:receiver,
		Kind:kind, Status:constant.NotificationStatusUnread, CreatedAt:time.Now()}
	if err = trx.Create(&notification).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func FindOnesUnreadNotificationsCount(uid int) (int, error) {
	var(
		err 	error
		count	int
	)
	if err = db.Model(&Notification{}).Where("receiver = ? AND status = ?", uid, constant.NotificationStatusUnread).
		Count(&count).Error; err != nil {
			log.Printf("获取未读消息数量出错：%+v\n", err)
			return count, err
	}
	return count, nil
}

func FindOnesNotifications(uid, pageIdx, pageSize int) ([]Notification, error) {
	var(
		err				error
		notifications	[]Notification
	)
	if err = db.Where("receiver = ? AND status != ?", uid, constant.NotificationStatusDeleted).
		Order("created_at").Offset(pageIdx*pageSize).Limit(pageSize).Find(&notifications).Error; err != nil {
		log.Printf("获取通知错误：%+v\n", err)
		return nil, err
	}
	return notifications, nil
}

func UpdateNotificationStatus(uid int, nidList []int, status int) error {
	var (
		err				error
		notifications	[]Notification
		targetNidList	[]int
	)
	trx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	trx.Set("gorm:query_option", "FOR UPDATE").Where("id in (?) AND uid = ?", nidList, uid).
		Find(&notifications)
	if len(notifications) == 0 {
		trx.Rollback()
		return nil
	}
	for _, item := range notifications {
		if item.Receiver == uid {
			targetNidList = append(targetNidList, item.Id)
		}
	}
	if err = db.Table("notification").Where("id IN (?)", targetNidList).
		Updates(map[string]interface{}{"status": status}).Error; err != nil {
		trx.Rollback()
		return err
	}
	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return err
	}
	return nil

}