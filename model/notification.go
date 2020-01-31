package model

type Notification struct {
	Id 			int		`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Pid			int		`json:"pid" gorm:"type:int"`
	Notifier	int		`json:"notifier" gorm:"type:int"`
	Receiver	int		`json:"receiver" gorm:"type:int"`
	Kind		int		`json:"kind" gorm:"type:int"`
	Status		int		`json:"status" gorm:"type:int"`
}
