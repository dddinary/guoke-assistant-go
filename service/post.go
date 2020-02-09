package service

import (
	"guoke-helper-golang/config"
	"guoke-helper-golang/model"
	"guoke-helper-golang/utils"
)

var pageSize int

func init() {
	pageSize = config.AppConf.PageSize
}

func GetNews(uid, kind, order, pageIdx int) (map[string]interface{}, error) {
	posts, err := model.FindPostsByCondition(kind, order, pageIdx, pageSize)
	if err != nil {
		return nil, err
	}
	return postsToRespMap(uid, posts), nil
}

func GetUserPost(uid, wantedUid, pageIdx int) (map[string]interface{}, error) {

	posts, err := model.FindPostsByUid(wantedUid, pageIdx, pageSize)
	if err != nil {
		return nil, err
	}
	return postsToRespMap(uid, posts), nil
}

func GetStaredPost(uid, pageIdx int) (map[string]interface{}, error) {
	posts, err := model.FindStaredPosts(uid, pageIdx, pageSize)
	if err != nil {
		return nil, err
	}
	return postsToRespMap(uid, posts), nil
}

func postsToRespMap(uid int, posts []model.Post) map[string]interface{} {
	var (
		stuInfoMap		map[int]interface{}
		postMapSlice	[]map[string]interface{}
		res				map[string]interface{}
	)
	stuInfoMap		= make(map[int]interface{})
	postMapSlice	= []map[string]interface{}{}
	res				= make(map[string]interface{})

	for _, post := range posts {
		postMap := utils.StructToMap(&post)
		if uid == 0 {
			postMap["liked"] = false
			postMap["stared"] = false
		} else {
			postMap["liked"] = model.IfLikedPost(uid, post.Id)
			postMap["stared"] = model.IfStared(uid, post.Id)
		}
		if post.Uid != 0 && post.Kind != 1{
			_, ok := stuInfoMap[post.Uid]
			if !ok {
				stu, _ := model.FindStudentById(post.Uid)
				stuInfoMap[stu.Id] = map[string]interface{}{"name": stu.Name, "avatar": stu.Avatar, "dpt": stu.Dpt}
			}
		}
		postMapSlice = append(postMapSlice, postMap)
	}
	res["students"] = stuInfoMap
	res["posts"] = postMapSlice
	return res

}

func GetPostDetail(uid, pid int) (map[string]interface{}, error) {
	var (
		err				error
		post			*model.Post
		stuInfoMap		map[int]interface{}
		commentSlice	[]map[string]interface{}
		res				map[string]interface{}
	)
	stuInfoMap		= make(map[int]interface{})
	commentSlice	= []map[string]interface{}{}
	res				= make(map[string]interface{})

	post, err = model.FindPostById(pid)
	if err != nil {
		return nil, err
	}
	postMap := utils.StructToMap(post)
	postMap["liked"] = model.IfLikedPost(uid, post.Id)
	postMap["stared"] = model.IfStared(uid, post.Id)
	stu, err := model.FindStudentById(post.Uid)
	if err == nil && stu != nil {
		stuInfoMap[stu.Id] = map[string]interface{}{"name":stu.Name, "dpt": stu.Dpt, "avatar": stu.Avatar, "status": stu.Status}
	}
	comments, err := model.FindCommentsByPostId(post.Id)
	if err != nil {
		return nil, err
	}
	for _, comment := range comments {
		commentMap := utils.StructToMap(&comment)
		if comment.Cid == 0 {
			commentMap["liked"] = false
		} else {
			commentMap["liked"] = model.IfLikedComment(uid, comment.Id)
		}
		commentSlice = append(commentSlice, commentMap)
		_, ok := stuInfoMap[comment.Uid]
		if !ok {
			stu, _ = model.FindStudentById(comment.Uid)
			stuInfoMap[stu.Id] = map[string]interface{}{"name": stu.Name, "avatar": stu.Avatar, "dpt": stu.Dpt}
		}
	}
	res["post"] = postMap
	res["comments"] = commentSlice
	res["students"] = stuInfoMap
	return res, nil
}

func AddPost(uid int, content string, kind int) error {
	return model.AddPost(uid, content, kind)
}

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

func DeletePost(uid, pid int) error {
	return model.DeletePost(uid, pid)
}
