package model

type Notification struct {
	Id 			int32		`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Pid			int32		`json:"pid" gorm:"type:int"`
	Notifier	int32		`json:"notifier" gorm:"type:int"`
	Receiver	int32		`json:"receiver" gorm:"type:int"`
	Kind		int32		`json:"kind" gorm:"type:int"`
	Status		int32		`json:"status" gorm:"type:int"`
}
