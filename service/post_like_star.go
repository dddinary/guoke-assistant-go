package service

import (
	"guoke-helper-golang/model"
	"log"
)

func LikePost(uid, pid int) error {
	return model.AddPostLike(uid, pid)
}

func UnlikePost(uid, pid int) error {
	return model.DeletePostLike(uid, pid)
}

func StarPost(uid, pid int) error {
	return model.AddStar(uid, pid)
}

func UnstarPost(uid, pid int) error {
	return model.DeleteStar(uid, pid)
}

func GetLikesByPostId(pid int) (map[int]interface{}, error) {
	var(
		err           error
		neededUidList []int
		stuInfoMap    map[int]interface{}
	)
	postLikes, err := model.FindPostLikesByPostId(pid)
	if err != nil {
		log.Printf("获取like列表出错: %+v\n", err)
		return nil, err
	}
	for _, postLike := range postLikes {
		neededUidList = append(neededUidList, postLike.Uid)
	}
	stuInfoMap, err = GetStudentsByIdList(neededUidList)
	if err != nil {
		log.Printf("获取like列表出错: %+v\n", err)
		return nil, err
	}
	return stuInfoMap, nil
}