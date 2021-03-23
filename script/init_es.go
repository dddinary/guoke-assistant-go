package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"time"
)


const indexName = "ucas_post"

var esClient	*elasticsearch.Client
var db *gorm.DB
var err error

type Post struct {
	Id 			int			`json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Uid			int			`json:"uid" gorm:"type:int"`
	Content		string		`json:"content" gorm:"type:text"`
	Kind		int			`json:"kind" gorm:"type:int"`
	Like		int			`json:"like" gorm:"type:int"`
	View		int			`json:"view" gorm:"type:int"`
	Comment		int			`json:"comment" gorm:"type:int"`
	CreatedAt	time.Time	`json:"created_at" gorm:"type:datetime"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"type:datetime"`
	Deleted		int			`json:"deleted" gorm:"type:int"`
}

func main() {
	initDBandES()

	var posts []Post
	err = db.Find(&posts).Error
	if err != nil {
		log.Panicf("获取posts列表失败：%v\n", err)
	}
	for _, post := range posts {
		fmt.Println(post.Content)
		//err = AddPostToES(post.Id, post.Uid, post.Content, post.CreatedAt, post.Deleted)
		//if err != nil {
		//	log.Printf("把post加入Elasticsearch出错：%v\n", err)
		//}
	}
	CloseDB()
}

func CloseDB() {
	db.Close()
}

func initDBandES() {

	dbType	:= "mysql"
	format	:= "%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local"
	connStr := fmt.Sprintf(format, "guoke", "Thisis4Guoke.", "localhost", "3306", "guoke_assistant", "utf8mb4")

	db, err = gorm.Open(dbType, connStr)
	if err != nil {
		log.Fatalln("Fail to connect database!")
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://127.0.0.1:9200",
		},
	}

	esClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("获取Elasticsearch客户端失败：%v\n", err)
	}
}

func AddPostToES(pid, uid int, content string, createdAt time.Time, deleted int) error {
	createAtStr := createdAt.Format("2006-01-02 15:04:05")
	post := map[string]interface{}{
		"id": pid,
		"uid": uid,
		"content": content,
		"created_at": createAtStr,
		"deleted": deleted,
	}
	payload, err := json.Marshal(post)
	if err != nil {
		return err
	}

	ctx := context.Background()
	res, err := esapi.CreateRequest{
		Index:      indexName,
		DocumentID: strconv.Itoa(pid),
		Body:       bytes.NewReader(payload),
	}.Do(ctx, esClient)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		}
		return fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
	}

	return nil
}
