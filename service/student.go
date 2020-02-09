package service

import (
	"guoke-helper-golang/constant"
	"guoke-helper-golang/model"
)

func BlockStudent(uid int) error {
	return model.UpdateStudentBlockStatus(uid, constant.StudentStatusBlocked)
}

func UnblockStudent(uid int) error {
	return model.UpdateStudentBlockStatus(uid, constant.StudentStatusCommon)
}
