package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/olivere/elastic/v7"
)

var (
	client    *elastic.Client
	indexName = "iot-edgewize-*"
	indexTest = "test"
)

// InitES 初始化 es client
func InitES(log elastic.Logger, addr ...string) error {
	//connect es
	var err error
	client, err = elastic.NewClient(
		elastic.SetURL(addr...),
		elastic.SetSniff(false), // 允许您指定弹性是否应该定期检查集群（默认为真）
		elastic.SetHealthcheck(false),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
		elastic.SetTraceLog(log),
	)
	if err != nil {
		log.Printf("init ES client error:%s", err.Error())
		return err
	}
	return nil
}

func main() {
	// 初始化 es client
	err := InitES(new(tracelog), "http://192.168.14.146:9200")
	if err != nil {
		panic(err)
	}

	// search()
	bulk()
}

// search 查询
func search() {
	search := client.Search(indexName)

	// 实例化一个bool搜索器
	boolQ := elastic.NewBoolQuery()
	// boolQ.Must(elastic.NewMatchQuery("appID", "EdgeMonitor"))  // 一级类目必须是鞋类
	// boolQ.Filter(elastic.NewRangeQuery("statusCode").Gte(925)) // 销量大于0
	// 打印查询语句
	resp, err := search.Query(boolQ).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Hits.Hits)

	// 组装查询
	res, err := search.Query(boolQ).From(0).Size(2).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("=====", *res.Hits.Hits[0])
}

// bulk 批量操作
func bulk() {
	// 批量插入
	// subjects := []Subject{
	// 	Subject{
	// 		ID:     1,
	// 		Title:  "肖恩克的救赎",
	// 		Genres: []string{"犯罪", "剧情"},
	// 	},
	// 	Subject{
	// 		ID:     2,
	// 		Title:  "千与千寻",
	// 		Genres: []string{"剧情", "喜剧", "爱情", "战争"},
	// 	},
	// }
	// bulkService := client.Bulk()
	// for _, subject := range subjects {
	// 	doc := elastic.NewBulkIndexRequest().Index(indexName).Id(strconv.Itoa(subject.ID)).Doc(subject)
	// 	bulkService = bulkService.Add(doc)

	// }
	// response, err := bulkService.Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// failed := response.Failed()
	// l := len(failed)
	// if l > 0 {
	// 	fmt.Printf("Error(%t)", response.Errors)
	// }

	// 删除、更新、插入在一个 bulk 中进行
	subject3 := Subject{
		ID:     3,
		Title:  "这个杀手太冷",
		Genres: []string{"剧情", "动作", "犯罪"},
	}
	subject4 := Subject{
		ID:     4,
		Title:  "阿甘正传",
		Genres: []string{"剧情", "爱情"},
	}

	subject5 := subject3
	subject5.Title = "这个杀手不太冷"

	index1Req := elastic.NewBulkIndexRequest().Index(indexTest).Id("3").Doc(subject3)
	index2Req := elastic.NewBulkIndexRequest().OpType("create").Index(indexTest).Id("4").Doc(subject4)
	delete1Req := elastic.NewBulkDeleteRequest().Index(indexTest).Id("1")
	update2Req := elastic.NewBulkUpdateRequest().Index(indexTest).Id("3").Doc(subject5)

	bulkRequest := client.Bulk()
	bulkRequest = bulkRequest.Add(index1Req)
	bulkRequest = bulkRequest.Add(index2Req)
	bulkRequest = bulkRequest.Add(delete1Req)
	bulkRequest = bulkRequest.Add(update2Req)

	_, err := bulkRequest.Refresh("wait_for").Do(context.Background())
	if err != nil {
		panic(err)
	}

	if bulkRequest.NumberOfActions() == 0 {
		fmt.Println("Actions all clear!")
	}

	searchResult, err := client.Search().
		Index(indexTest).
		Sort("id", false). // 按id升序排序
		Pretty(true).
		Do(context.Background()) // 执行
	if err != nil {
		panic(err)
	}
	var subject Subject
	for _, item := range searchResult.Each(reflect.TypeOf(subject)) {
		if t, ok := item.(Subject); ok {
			fmt.Printf("Found: Subject(id=%d, title=%s)\n", t.ID, t.Title)
		}
	}
}

type Subject struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Genres []string `json:"genres"`
}
