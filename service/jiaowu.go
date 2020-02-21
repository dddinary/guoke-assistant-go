package service

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
	"guoke-assistant-go/config"
	"guoke-assistant-go/model"
	"guoke-assistant-go/utils"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const baseURL		= "http://sep.ucas.ac.cn"
const siteJiaoWu	= "jwxk"
const siteCourse	= "course"
const siteEPay		= "epay"

const scienceLecture	= "http://jwxk.ucas.ac.cn/subject/lecture"
const humanityLecture	= "http://jwxk.ucas.ac.cn/subject/humanityLecture"

type site struct {
	id			int
	url			string
	loginUrl	string
	roleId		int
}

var sites map[string]site
var userClients map[string]*req.Req

func init() {
	userClients = make(map[string]*req.Req)
	sites = make(map[string]site)
	sites[siteJiaoWu] = site{
		id: 226,
		url: "http://jwxk.ucas.ac.cn",
		loginUrl: "http://jwxk.ucas.ac.cn/login",
		roleId: 821,
	}
	sites[siteCourse] = site{
		id: 16,
		url: "http://course.ucas.ac.cn",
		loginUrl: "http://course.ucas.ac.cn/portal/plogin",
		roleId: 801,
	}
	sites[siteEPay] = site {
		id: 311,
		url: "http://epay.ucas.ac.cn",
		loginUrl: "http://epay.ucas.ac.cn/NetWorkUI/sepLogin.htm",
		roleId: 1800,
	}
}

// 因为不需要验证码，所以也就不需要openid标识不同的用户了，直接使用一次性的 client 就行
func openidToClient(openid string) *req.Req {
	cli, ok := userClients[openid]
	if !ok || cli == nil{
		userClients[openid] = req.New()
		cli = userClients[openid]
		time.AfterFunc(10 * time.Minute, func() {
			delete(userClients, openid)
		})
	}
	return cli
}

func GetCaptcha(openid string) (img []byte) {
	cli := openidToClient(openid)
	picUrl := baseURL + "/changePic"
	resp, err := cli.Get(picUrl)
	if err != nil || resp.Response().StatusCode != http.StatusOK {
		log.Printf("获取验证码出错：%v\n", err)
		return nil
	}
	img = resp.Bytes()
	return
}

func LoginAndGetCourse(openid, username, pwd, avatar string) map[string]interface{} {
	var name, dpt, token string
	var uid int
	// cli := openidToClient(openid)
	cli := req.New()
	if !MainLoginWithoutCaptcha(cli, username, pwd) {
		log.Printf("登录失败")
		return nil
	}
	stu, _ := model.FindStudentByAccount(username)
	nameDpt := make(map[string]string)

	if stu == nil {
		nameDpt = findNameAndDpt(cli)
		name = nameDpt["name"]
		dpt = nameDpt["dpt"]
		if avatar == "" {
			avatar  = utils.BTGetAvatarUrl()
		}
		token, uid, _ = model.AddStudent(username, name, dpt, avatar, openid)
	} else {
		uid = stu.Id
		name = stu.Name
		dpt = stu.Dpt
		avatar = stu.Avatar
		token = stu.UpdateToken()
	}

	if !siteLogin(cli, siteJiaoWu) {
		log.Printf("登录站点失败")
		return nil
	}
	cidList := getCourseList(cli)
	courseDetail, timeTable := GetCourseDetailAndTimeTable(cidList)
	return map[string]interface{}{
		"id": uid,
		"name": name,
		"dpt": dpt,
		"avatar": avatar,
		"token": token,
		"courseList": cidList,
		"courseDetail": courseDetail,
		"timeTable": timeTable,
	}
}

type MainLoginRes struct {
	F	bool	`json:"f"`
	Msg	string	`json:"msg"`
}

func MainLoginWithoutCaptcha(cli *req.Req, username, password string) bool {
	var (
		err			error
		resp		*req.Resp
		loginRes	MainLoginRes
		errLogMsg	= "登录失败"
	)
	_, _ = cli.Get("http://onestop.ucas.edu.cn")
	data := req.Param{
		"username": username,
		"password": password,
	}
	headers := req.Header{
		"X-Requested-With": "XMLHttpRequest",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
	}
	loginUrl := "http://onestop.ucas.edu.cn/Ajax/Login/0"
	resp, err = cli.Post(loginUrl, data, headers)
	if err != nil {
		log.Printf("%s：%+v\n", errLogMsg, err)
		return false
	}
	if err = resp.ToJSON(&loginRes); err != nil {
		log.Printf("%s：%+v\n", errLogMsg, err)
		return false
	}
	if !loginRes.F {
		log.Printf("%s：%s", errLogMsg, loginRes.Msg)
		return false
	}
	resp, err = cli.Get(loginRes.Msg)
	if err != nil {
		log.Printf("%s：%+v\n", errLogMsg, err)
		return false
	}
	return true
}

func MainLoginWithCaptcha(cli *req.Req, username, pwd, cap string) bool {
	data := req.Param{
		"userName": username,
		"pwd": pwd,
		"certCode": cap,
		"sb": "sb",
	}
	loginUrl := baseURL + "/slogin"
	resp, err := cli.Post(loginUrl, data)
	if err != nil {
		log.Printf("登录失败：%+v\n", err)
		return false
	}
	redirectPath := resp.Response().Request.URL.Path
	if !strings.Contains(redirectPath, "appStore") {
		log.Printf("登录失败, redirect to：%s\n", redirectPath)
		return false
	}
	return true
}

func findNameAndDpt(cli *req.Req) map[string]string {
	resp, err := cli.Get(baseURL + "/appStore")
	if err != nil {
		log.Printf("登录失败：%+v\n", err)
		return nil
	}
	body := resp.String()
	body = strings.ReplaceAll(body, "\n", "")
	re := regexp.MustCompile(`当前用户所在单位">\s*(.*?)</li>`)
	match := re.FindAllStringSubmatch(body, 1)
	if len(match) < 1 || len(match[0]) < 2 {
		log.Printf("获取姓名单位错误：正则表达式没找到正确匹配")
		return nil
	}
	temp := strings.ReplaceAll(match[0][1], " ", "")
	temp = strings.ReplaceAll(temp, "\r", "")
	temp = strings.ReplaceAll(temp, "\t", "")
	nameDpt := strings.Split(temp, "&nbsp;")
	if len(nameDpt) != 2 {
		log.Printf("获取姓名单位错误：split后不是两个字符串")
		return nil
	}
	return map[string]string{"name": nameDpt[1], "dpt": nameDpt[0]}
}

func siteLogin(cli *req.Req, siteName string) bool {
	targetSite := sites[siteName]
	idGetUrl := baseURL + "/portal/site/" + strconv.Itoa(targetSite.id)
	resp, err := cli.Get(idGetUrl)
	if err != nil {
		log.Printf("登录站点%s失败：%v", siteName, err)
		return false
	}
	body := resp.String()
	re := regexp.MustCompile(`Identity=([-0-9a-z]+)&`)
	match := re.FindAllStringSubmatch(body, 1)
	if len(match) < 1 || len(match[0]) < 2 {
		log.Printf("登录站点%s失败：找不到Identity", siteName)
		return false
	}
	identity := match[0][1]
	siteUrl := targetSite.loginUrl + "?Identity=" + identity + "&roleId=" + strconv.Itoa(int(targetSite.roleId))
	_, err = cli.Get(siteUrl)
	if err != nil {
		log.Printf("登录站点%s失败：%v", siteName, err)
		return false
	}
	return true
}

func getCourseList(cli *req.Req) []int {
	resp, err := cli.Get(sites[siteJiaoWu].url + "/courseManage/selectedCourse")
	if err != nil {
		log.Printf("获取课程列表失败：%v", err)
		return nil
	}
	body := resp.String()
	body = strings.ReplaceAll(body, "\n", "")
	re := regexp.MustCompile(`/courseplan/(\d+)"(.*?)第(.*?)学期`)
	match := re.FindAllStringSubmatch(body, -1)
	if len(match) < 1 || len(match[0]) < 2 {
		return nil
	}
	var cidList []int
	for _, item := range match {
		if item[3] == "二" {
			cid, err := strconv.ParseInt(item[1], 10, 32)
			if err != nil {
				continue
			}
			cidList = append(cidList, int(cid))
		}
	}
	return cidList
}

func GetCourseDetailAndTimeTable(cidList []int) (map[int]interface{}, [21][8][]interface{}) {
	var table [21][8][]interface{}
	for i := range table {
		for j := range table[i] {
			table[i][j] = []interface{}{}
		}
	}
	courseDetail := make(map[int]interface{})
	courses, _ := model.FindCoursesByCidList(cidList)
	for _, course := range courses {
		courseMap := utils.StructToMap(&course)
		timePlace := model.FindTimePlaceByCid(course.Cid)
		courseMap["time_place"] = timePlace
		courseDetail[course.Cid] = courseMap
		for _, tp := range timePlace {
			weekDay := tp.Weekday
			weekNoList := strings.Split(tp.Weekno, ",")
			for _, wn := range weekNoList {
				weekNo, _ := strconv.Atoi(wn)
				table[weekNo][weekDay] = append(table[weekNo][weekDay], tp)
			}
		}
	}
	return courseDetail, table
}

func UpdateLectureList() error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%+v\n", r)
		}
	}()
	var (
		err		error
	)
	cli := req.New()
	if !MainLoginWithoutCaptcha(cli, config.AdminConf.Account, config.AdminConf.Pwd) {
		err = errors.New("获取讲座信息时登录失败")
		log.Printf("更新讲座信息出错：%v\n", err)
		return err
	}
	if !siteLogin(cli, siteJiaoWu) {
		err = errors.New("获取讲座信息时登录站点失败")
		log.Printf("更新讲座信息出错：%v\n", err)
		return err
	}
	humanityLidList := make([]int, 0)
	scienceLidList := make([]int, 0)
	humanityLidList = append(humanityLidList, GetPageList(cli, humanityLecture, 1)...)
	for i := 1; i <= 3; i++ {
		scienceLidList = append(scienceLidList, GetPageList(cli, scienceLecture, i)...)
	}
	addLectures(cli, humanityLidList, 2)
	addLectures(cli, scienceLidList, 1)
	return nil
}

func GetPageList(cli *req.Req, reqUrl string, pageNum int) []int {
	var (
		err			error
		resp		*req.Resp
		lidList 	= make([]int, 0)
	)
	param := req.Param{
		"pageNum":  pageNum,
	}
	resp, err = cli.Post(reqUrl, param)
	if err != nil {
		log.Printf("获取讲座列表出错：%v\n", err)
		return lidList
	}
	body := resp.String()
	re := regexp.MustCompile(`/subject/(\d+)/view`)
	match := re.FindAllStringSubmatch(body, -1)
	for _, matchList := range match {
		lid, _ := strconv.Atoi(matchList[1])
		lidList = append(lidList, lid)
	}
	return lidList
}

func addLectures(cli *req.Req, lidList []int, category int) {
	var (
		err			error
		hasIt		bool
		resp		*req.Resp
		matched		[][]string
		re			*regexp.Regexp
		name, dpt, venue1, venue2, venue, desc, pic string
		startTime, endTime time.Time
		timeLayout	= "2006/01/02 15:04:05"
	)
	urls := map[int]string{
		1: "http://jwxk.ucas.ac.cn/subject/%d/view",
		2: "http://jwxk.ucas.ac.cn/subject/%d/humanityView",
	}
	urlPattern := urls[category]
	for _, lid := range lidList {
		hasIt, err = LectureExists(lid)
		if hasIt {
			log.Printf("该记录已存在")
			continue
		}
		url := fmt.Sprintf(urlPattern, lid)
		resp, err = cli.Get(url)
		if err != nil {
			log.Printf("获取讲座信息出错，lid:%d, %+v\n", lid, err)
			continue
		}
		body := resp.String()
		body = strings.ReplaceAll(body, "\n", "")
		re = regexp.MustCompile(`讲座名称：\s*(.*?)</td>`)
		matched = re.FindAllStringSubmatch(body, 1)
		if len(matched) < 1 || len(matched[0]) < 2 {
			continue
		}
		name = matched[0][1]
		re = regexp.MustCompile(`部门：\s*(.*?)</td>`)
		matched = re.FindAllStringSubmatch(body, 1)
		if len(matched) < 1 || len(matched[0]) < 2 {
			continue
		}
		dpt = matched[0][1]
		re = regexp.MustCompile(`开始时间：\s*(.*?)</td>`)
		matched = re.FindAllStringSubmatch(body, 1)
		if len(matched) < 1 || len(matched[0]) < 2 {
			continue
		}
		startTime, _ = time.ParseInLocation(timeLayout, matched[0][1], time.Local)
		re = regexp.MustCompile(`结束时间：\s*(.*?)</td>`)
		matched = re.FindAllStringSubmatch(body, 1)
		if len(matched) < 1 || len(matched[0]) < 2 {
			continue
		}
		endTime, _ = time.ParseInLocation(timeLayout, matched[0][1], time.Local)
		re = regexp.MustCompile(`主会场地点：\s*(.*?)</td>`)
		matched = re.FindAllStringSubmatch(body, 1)
		if len(matched) < 1 || len(matched[0]) < 2 {
			continue
		}
		venue1 = matched[0][1]
		re = regexp.MustCompile(`分会场地点：\s*(.*?)</td>`)
		matched = re.FindAllStringSubmatch(body, 1)
		if len(matched) < 1 || len(matched[0]) < 2 {
			continue
		}
		venue2 = matched[0][1]
		venue = venue1 + "#" + venue2
		re = regexp.MustCompile(`讲座介绍</td>\s*</tr>\s*<tr>\s*<td colspan="2">(.*?)</td>`)
		matched = re.FindAllStringSubmatch(body, 1)
		if len(matched) < 1 || len(matched[0]) < 2 {
			continue
		}
		desc = matched[0][1]
		pic = utils.BTGetWallPaperUrl()
		log.Println(name, dpt, startTime, endTime, venue, pic)
		err = AddLecture(lid, name, category, dpt, startTime, endTime, venue, desc, pic)
		if err != nil {
			log.Printf("增加讲座信息出错：%+v\n", err)
		}
	}
}
