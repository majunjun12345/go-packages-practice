package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	Marshal：将 go 对象转换为 json 对象
	UnMarshal：将 json 对象转换为为 go 对象，作用于内存中的 []byte

	会自动进行列表转换
	作用于导出字段，忽略非导出字段
	可以通过 tag 自定义字段名，不用 tag 则默认

	decoder：从 reader 流 或 解析大数据，或 resp body，参数是 reader
		json.Decoder会一个一个元素进行加载，不会把整个json数组读到内存里面
		如果使用 token(必须是数组)，需要建立 for 循环，读出的是 struct 对象
		如果不使用 token，不需要 for 循环，读出的是 struct 对象数组

	json 和 map 也可以通过该包进行转换
	json 转 map 的适合，虽然 map 本身就是引用类型，但还是要加上 & 取址

	map 与 struct 互转需要用到第三方库 mapstructure
*/

type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
	desc   string   `json:desc`
}

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

type Message struct {
	Name string
	Text string
}

func main() {
	// commonMarshal()
	// commonUnMarshal()
	// decoder()
	json2Map()
	map2Json()
}

func commonMarshal() {
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
		desc:   "good",
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
}

// json 数组可以直接处理
func commonUnMarshal() {
	var m []Message
	err := json.Unmarshal([]byte(jsonStream0), &m)
	fmt.Println(err)
	fmt.Printf("%+v\n", m)
}

/*
	decoder
*/
func decoder() {
	//  json 数组
	newReader := strings.NewReader(jsonStream0)
	decode := json.NewDecoder(newReader)
	t, err := decode.Token()
	CheckErr(err)
	fmt.Println("json delim:", t) // [

	for decode.More() {
		var m Message // 变量在 for 循环里面
		err := decode.Decode(&m)
		CheckErr(err)
		fmt.Printf("%+v\n", m)
	}

	t, err = decode.Token()
	CheckErr(err)
	fmt.Println("json delim:", t) // ]

	// json 流
	newReader1 := json.NewDecoder(strings.NewReader(jsonStream1))
	for {
		var m1 Message // 变量在 for 循环里面
		err = newReader1.Decode(&m1)
		if err == io.EOF {
			break
		} else if err != nil {
			CheckErr(err)
		}
		fmt.Printf("%+v\n", m1)
	}

	// json 流 数组 不能这么搞
	newReader2 := json.NewDecoder(strings.NewReader(jsonStream1))
	var m2 []Message // 变量在 for 循环里面
	for {

		err = newReader2.Decode(&m2)
		if err == io.EOF {
			break
		} else if err != nil {
			CheckErr(err)
		}
	}
	fmt.Printf("%+v\n", m2)

	// file 中读取
	file, err := os.Open("test.json")
	defer file.Close()
	CheckErr(err)

	var f []*Feed
	json.NewDecoder(file).Decode(&f)
	fmt.Printf("this is m2:%+v\n", f)

	// r := json.NewDecoder(file)
	// tt, _ := r.Token()
	// fmt.Println(tt)
	// for r.More() {
	// 	var fs Feed
	// 	r.Decode(&fs)
	// 	fmt.Printf("this is m3:%+v\n", fs)
	// }
	// tt, _ = r.Token()
	// fmt.Println(tt)
}

const jsonStream0 = `
	[
		{"Name": "Ed", "Text": "Knock knock."},
		{"Name": "Sam", "Text": "Who's there?"},
		{"Name": "Ed", "Text": "Go fmt."},
		{"Name": "Sam", "Text": "Go fmt who?"},
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	]
`

const jsonStream1 = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// ----------------- json ---- map

func json2Map() {
	jsonStr := `
	{
			"name": "jqw",
			"age": 18
	}
`
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult) // 需要引用类型
	fmt.Println(err)
	fmt.Println(mapResult)
}

func map2Json() {
	mapInstances := []map[string]interface{}{}
	instance_1 := map[string]interface{}{"name": "John", "age": 10}
	instance_2 := map[string]interface{}{"name": "Alex", "age": 12}
	mapInstances = append(mapInstances, instance_1, instance_2)

	jsonData, err := json.Marshal(instance_1)
	fmt.Println(err)
	fmt.Println(jsonData)

	jsonDataList, err := json.Marshal(mapInstances)
	fmt.Println(err)
	fmt.Println(jsonDataList)
}
