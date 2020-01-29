package model

import "testing"

func TestAddGroup(t *testing.T) {
	t.Log(AddGroup("guokezhushou", "123456", "果壳助手", "https://tva3.sinaimg.cn/small/9bd9b167ly1g1p9phfd02j20b40b4t91.jpg"))
	t.Log(CheckGroupPwd("guokezhushou", "123456"))
	t.Log(CheckGroupPwd("guokezhushoutest", "123456"))
	t.Log(CheckGroupPwd("guokezhushou", "1234567"))
}
