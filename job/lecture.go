package job

import (
	"guoke-assistant-go/service"
	"guoke-assistant-go/utils"
	"log"
)

type LectureJob struct {

}

func (lecJob LectureJob) Run() {
	var (
		err		error
	)
	log.Println("开始更新讲座信息")
	utils.BotMsgJobStart("讲座信息爬取")
	err = service.UpdateLectureList()
	if err != nil {
		log.Printf("更新讲座列表出错：%+v\n", err)
	}
	err = service.DeleteLectureDataInRedis()
	if err != nil {
		log.Printf("Redis中删除讲座缓存出错：%+v\n", err)
	}
}
