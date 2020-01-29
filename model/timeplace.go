package model

type Timeplace struct {
	Id 		uint		`json:"id" gorm:"primary_key"`
	Cid		int32		`json:"cid" form:"cid"`
	WeekDay	int32		`json:"weekday" form:"weekday"`
	Jie		string		`json:"dpt" jie:"type:varchar(255)"`
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
	res["weekday"]	= tp.WeekDay
	res["jie"]		= tp.Jie
	res["room"]		= tp.Room
	res["weekno"]	= tp.Weekno
	return res
}