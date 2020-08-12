package service

import (
	"guoke-assistant-go/constant"
	"guoke-assistant-go/model"
	"guoke-assistant-go/utils"
)

var pageSize int

func init() {
	pageSize = constant.PageSize
}

func GetNews(uid, kind, order, pageIdx int) (map[string]interface{}, error) {
	posts, err := model.FindPostsByCondition(kind, order, pageIdx, pageSize)
	if err != nil {
		return nil, err
	}
	return postsToRespMap(uid, posts), nil
}

func GetUserPost(uid, wantedUid, pageIdx int) (map[string]interface{}, error) {
	var canSeePosts []model.Post
	posts, err := model.FindPostsByUid(wantedUid, pageIdx, pageSize)
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		if post.Kind != constant.PostKindAnonymous || uid == wantedUid {
			canSeePosts = append(canSeePosts, post)
		}
	}
	return postsToRespMap(uid, canSeePosts), nil
}

func GetStaredPost(uid, pageIdx int) (map[string]interface{}, error) {
	posts, err := model.FindStaredPosts(uid, pageIdx, pageSize)
	if err != nil {
		return nil, err
	}
	return postsToRespMap(uid, posts), nil
}

func SearchPost(words string, uid int, pageIdx int) (map[string]interface{}, error) {
	posts, err := model.FindPostsByWords(words, pageIdx, pageSize)
	if err != nil {
		return nil, err
	}
	return postsToRespMap(uid, posts), nil
}

func postsToRespMap(uid int, posts []model.Post) map[string]interface{} {
	var (
		neededUidList 	[]int
		stuInfoMap		map[int]interface{}
		postMapSlice	[]map[string]interface{}
		res				map[string]interface{}
	)
	postMapSlice	= []map[string]interface{}{}
	res				= make(map[string]interface{})

	for _, post := range posts {
		if post.Kind == constant.PostKindAnonymous {
			post.Uid = 0
		}
		if post.Deleted == 1 {
			continue
		}
		postMap := utils.StructToMap(&post)
		if uid == 0 {
			postMap["liked"] = false
			postMap["stared"] = false
		} else {
			postMap["liked"] = model.IfLikedPost(uid, post.Id)
			postMap["stared"] = model.IfStared(uid, post.Id)
		}
		images, err := model.FindImagesByPostId(post.Id)
		if err != nil {
			images = []string{}
		}
		postMap["imgCount"] = len(images)
		postMap["images"] = images
		if post.Uid != 0 && post.Kind != constant.PostKindAnonymous {
			neededUidList = append(neededUidList, post.Uid)
		}
		postMapSlice = append(postMapSlice, postMap)
	}
	stuInfoMap, _ = GetStudentsNoSecretInfoByIdList(neededUidList)
	res["students"] = stuInfoMap
	res["posts"] = postMapSlice
	return res
}

func getPostsContentAbstractionByPIdList(pidList []int) (map[int]interface{}, error) {
	res := make(map[int]interface{})
	posts, err := model.FindPostsByIdList(pidList)
	if err != nil {
		return res, err
	}
	var abstraction string
	for _, post := range posts {
		if len(post.Content) < 20 {
			abstraction = post.Content
		} else {
			abstraction = post.Content[:20]
		}
		res[post.Id] = abstraction
	}
	return res, nil
}

func GetPostDetail(uid, pid int) (map[string]interface{}, error) {
	var (
		err				error
		post			*model.Post
		stuInfoMap		map[string]interface{}
		res				map[string]interface{}
	)
	res				= make(map[string]interface{})

	post, err = model.FindPostById(pid)
	if err != nil || post == nil {
		return nil, err
	}
	if post.Deleted == 1 {
		return nil, model.ErrorPostNotFound
	}
	if post.Kind == constant.PostKindAnonymous {
		post.Uid = 0
	} else {
		stuInfoMap, _ = GetStudentNoSecretInfoById(post.Uid)
	}
	images, err := model.FindImagesByPostId(post.Id)
	if err != nil {
		images = []string{}
	}
	postMap := utils.StructToMap(post)
	postMap["liked"] = model.IfLikedPost(uid, post.Id)
	postMap["stared"] = model.IfStared(uid, post.Id)
	postMap["imgCount"] = len(images)
	postMap["images"] = images
	res["post"] = postMap
	res["student"] = stuInfoMap
	return res, nil
}

func AddPost(uid int, content string, kind int, images []string) error {
	return model.AddPost(uid, content, kind, images)
}

func DeletePost(uid, pid int) error {
	return model.DeletePost(uid, pid)
}
