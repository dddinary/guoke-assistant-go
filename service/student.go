package service

import (
	"guoke-helper-golang/constant"
	"guoke-helper-golang/model"
	"log"
	"strconv"
)

func BlockStudent(uid int) error {
	return model.UpdateStudentBlockStatus(uid, constant.StudentStatusBlocked)
}

func UnblockStudent(uid int) error {
	return model.UpdateStudentBlockStatus(uid, constant.StudentStatusCommon)
}

func stuModelToMap(student *model.Student) map[string]string {
	res := make(map[string]string)
	res["id"] = strconv.Itoa(student.Id)
	res["name"] = student.Name
	res["dpt"] = student.Dpt
	res["avatar"] = student.Avatar
	return res
}

func GetStudentNoSecretInfoById(sid int) (map[string]string, error) {
	student, err := model.FindStudentById(sid)
	if err != nil {
		log.Printf("获取学生信息出错：%+v\n", err)
		return nil, err
	}
	stuInfo := stuModelToMap(student)
	return stuInfo, nil
}

func GetStudentsNoSecretInfoByIdList(idList []int) (map[int]interface{}, error) {
	var(
		err		error
		res		map[int]interface{}
	)
	res = make(map[int]interface{})
	students, err := model.FindStudentsByIdList(idList)
	if err != nil {
		log.Printf("获取学生列表出错：%+v\n", err)
		return res, err
	}
	for _, stu := range students {
		res[stu.Id] = stuModelToMap(&stu)
	}
	return res, nil
}