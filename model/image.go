package model

type Image struct {
	Id 			int			`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Pid			int			`json:"pid" gorm:"type:int"`
	Url			string		`json:"content" gorm:"type:text"`
	Idx			int			`json:"idx" gorm:"type:int"`
}

func FindImagesByPostId(pid int) []string {
	var (
		urls	[]string
		images	[]Image
	)
	if err := db.Where("pid = ?", pid).Order("idx").Find(&images).Error; err != nil {
		return urls
	}
	for _, image := range images {
		urls = append(urls, image.Url)
	}
	return urls
}