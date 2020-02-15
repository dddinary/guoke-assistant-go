package service

import (
	"guoke-assistant-go/constant"
	"guoke-assistant-go/model"
)

func GetUnreadNotificationsCount(uid int) int {
	var (
		err 	error
		count	int
	)
	count, err = model.FindOnesUnreadNotificationsCount(uid)
	if err != nil {
		return 0
	}
	return count
}

func GetOnesNotifications(uid, pageIdx int) (map[string]interface{}, error) {
	var (
		err				error
		stuInfoMap		map[int]interface{}
		notifications	[]model.Notification
		res				map[string]interface{}
	)
	res				= make(map[string]interface{})
	notifications, err = model.FindOnesNotifications(uid, pageIdx, pageSize)
	if err != nil {
		return res, err
	}
	var neededUidList []int
	for _, item := range notifications {
		neededUidList = append(neededUidList, item.Notifier)
	}
	stuInfoMap, _ = GetStudentsNoSecretInfoByIdList(neededUidList)
	res["students"] = stuInfoMap
	res["notifications"] = notifications
	return res, nil
}

func MarkReadNotifications(uid int, nidList []int) error {
	return model.UpdateNotificationStatus(uid, nidList, constant.NotificationStatusRead)
}

func DeleteNotifications(uid int, nidList []int) error {
	return model.UpdateNotificationStatus(uid, nidList, constant.NotificationStatusDeleted)
}
