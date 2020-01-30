package model

type Timeplace struct {
	Id 		int32		`json:"id" gorm:"primary_key"`
	Cid		int32		`json:"cid" gorm:"type:int"`
	Weekday	int32		`json:"weekday" gorm:"type:int"`
	Jie		string		`json:"jie" gorm:"type:varchar(255)"`
	Room	string		`json:"room" gorm:"type:varchar(255)"`
	Weekno	string		`json:"weekno" gorm:"type:varchar(255)"`
}

func FindTimePlaceByCid(cid int32) []Timeplace {
	var timeplaces []Timeplace
	db.Where("cid = ?", cid).Find(&timeplaces)
	return timeplaces
}

func (tp *Timeplace)ToMap() map[string]interface{} {
	res := make(map[string]interface{})
	res["id"]		= tp.Id
	res["cid"]		= tp.Cid
	res["weekday"]	= tp.Weekday
	res["jie"]		= tp.Jie
	res["room"]		= tp.Room
	res["weekno"]	= tp.Weekno
	return res
}