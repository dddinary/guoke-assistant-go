package model

type Image struct {
	Id 			uint		`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Pid			int32		`json:"uid" gorm:"type:int"`
	Url			string		`json:"content" gorm:"type:text"`
	Idx			int32		`json:"kind" gorm:"type:int"`
}
