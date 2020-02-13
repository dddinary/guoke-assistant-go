package service

import (
	"guoke-helper-golang/model"
	"guoke-helper-golang/utils"
	"log"
)

func CommentPost(uid, pid int, content string) error {
	return model.AddComment(uid, pid, 0, content)
}

func CommentComment(uid, pid, cid int, content string) error {
	return model.AddComment(uid, pid, cid, content)
}

func LikeComment(uid, cid int) error {
	return model.AddCommentLike(uid, cid)
}

func UnlikeComment(uid, cid int) error {
	return model.DeleteCommentLike(uid, cid)
}

func DeleteComment(uid, cid int) error {
	return model.DeleteComment(uid, cid)
}

func GetCommentsByPostId(uid, pid int) (map[string]interface{}, error) {
	var (
		err				error
		stuInfoMap		map[int]interface{}
		commentSlice	[]map[string]interface{}
		res				map[string]interface{}
	)
	commentSlice	= []map[string]interface{}{}
	res				= make(map[string]interface{})

	comments, err := model.FindCommentsByPostId(pid)
	if err != nil {
		log.Printf("获取评论出错：%+v\n", err)
		return nil, err
	}
	var neededUidList []int
	for _, comment := range comments {
		commentMap := utils.StructToMap(&comment)
		if comment.Cid == 0 || uid == 0 {
			commentMap["liked"] = false
		} else {
			commentMap["liked"] = model.IfLikedComment(uid, comment.Id)
		}
		commentSlice = append(commentSlice, commentMap)
		_, ok := stuInfoMap[comment.Uid]
		if !ok {
			neededUidList = append(neededUidList, comment.Uid)
		}
		stuInfoMap, _ = GetStudentsNoSecretInfoByIdList(neededUidList)
	}
	res["comments"] = commentSlice
	res["students"] = stuInfoMap
	return res, nil
}
