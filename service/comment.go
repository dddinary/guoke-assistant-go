package service

import "guoke-helper-golang/model"

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
