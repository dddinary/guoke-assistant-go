package model

import "time"

type Comment struct {
	Id 			uint		`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Pid			int32		`json:"pid" gorm:"type:int"`
	Uid			int32		`json:"uid" gorm:"type:int"`
	Content		string		`json:"content" gorm:"type:text"`
	Kind		int32		`json:"kind" gorm:"type:int"`
	Like		int32		`json:"like" gorm:"type:int"`
	CreatedAt	time.Time	`json:"created_at" gorm:"type:datetime"`
	Deleted		int32		`json:"deleted" gorm:"type:int"`
}
