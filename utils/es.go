package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"guoke-assistant-go/config"
	"log"
	"strconv"
	"strings"
	"time"
)

const indexName = "ucas_post"

var esClient	*elasticsearch.Client

func init() {
	esAddress := config.ESConf.Address
	cfg := elasticsearch.Config{
		Addresses: []string{
			esAddress,
		},
	}
	var err error
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

func MarkPostInESDeleted(pid int) error {
	payload := `{"doc": {"deleted": 1}}`
	res, err := esClient.Update(indexName, strconv.Itoa(pid), strings.NewReader(payload))
	if err != nil {
		return err
	}
	log.Printf("%v\n", res)
	return nil
}

func SearchPostInES(words string, pageIdx, pageSize int) (pidList []int, err error) {
	pidList = []int{}
	var r map[string]interface{}
	from := pageIdx * pageSize
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"content": words,
						},
					},
				},
				"filter": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"deleted": 0,
						},
					},
				},
			},
		},
		"from": from,
		"size": pageSize,
	}
	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s", err)
		return
	}

	// Perform the search request.
	var res *esapi.Response
	res, err = esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex("ucas_post"),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err = json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %s", err)
			return
		} else {
			// Print the response status and error information.
			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
			return
		}
	}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return
	}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		pidList = append(pidList, int(hit.(map[string]interface{})["_source"].(map[string]interface{})["id"].(float64)))
	}
	return
}
