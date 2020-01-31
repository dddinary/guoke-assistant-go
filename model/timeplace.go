package model

type Timeplace struct {
	Id 		int			`json:"id" gorm:"primary_key"`
	Cid		int			`json:"cid" gorm:"type:int"`
	Weekday	int			`json:"weekday" gorm:"type:int"`
	Jie		string		`json:"jie" gorm:"type:varchar(255)"`
	Room	string		`json:"room" gorm:"type:varchar(255)"`
	Weekno	string		`json:"weekno" gorm:"type:varchar(255)"`
}

func FindTimePlaceByCid(cid int) []Timeplace {
	var timeplaces []Timeplace
	db.Where("cid = ?", cid).Find(&timeplaces)
	return timeplaces
}
