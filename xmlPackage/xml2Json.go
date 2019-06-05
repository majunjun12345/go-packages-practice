package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type DataFormat struct {
	ProductList []struct { // 结构体中嵌套结构体
		Sku    string `xml:"sku" json:"sku"`
		Quanty int    `xml:"quanty" json:"quanty"`
	} `xml:"Product" json:"products"`
}

func main() {
	/*
		xml 先转 struct 再 转 json
	*/
	// xml2Json()

	/*
		xml 文件转 json 文件
	*/

	type Staff struct {
		Id        int    `xml:"id"`
		FirstNmae string `xml:"firstname"`
		LastName  string `xml:"lastname"`
		UserName  string `xml:"username"`
	}

	type Company struct {
		xmlName xml.Name `xml:company`
		Staffs  []Staff  `xml:"staff"`
	}

	xmlFile, err := ioutil.ReadFile("Employees.xml")
	if err != nil {
		panic(err)
	}

	c := &Company{}
	err = xml.Unmarshal(xmlFile, c)
	if err != nil {
		panic(nil)
	}

	fmt.Println(c.Staffs)

	var oneStaff Staff
	var allStass []Staff

	for _, v := range c.Staffs {
		oneStaff.FirstNmae = v.FirstNmae
		oneStaff.Id = v.Id
		oneStaff.LastName = v.LastName
		oneStaff.UserName = v.UserName

		allStass = append(allStass, oneStaff)
	}
	jsonData, err := json.Marshal(allStass)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))

	jsonFile, err := os.Create("Employees.json")
	defer jsonFile.Close()
	if err != nil {
		panic(nil)
	}
	var b bytes.Buffer
	err = json.Indent(&b, jsonData, "", "\t") // json 的格式化输出
	b.WriteTo(jsonFile)
}

func xml2Json() {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8" ?>
					<ProductList>
						<Product>
							<sku>ABC123</sku>
							<quantity>2</quantity>
						</Product>
						<Product>
							<sku>ABC123</sku>
							<quantity>2</quantity>
						</Product>
					</ProductList>`)
	data := &DataFormat{}
	err := xml.Unmarshal(xmlData, data)
	if err != nil {
		panic(nil)
	}
	fmt.Printf("%+v\n", data)

	result, err := json.Marshal(data)
	if err != nil {
		panic(nil)
	}

	fmt.Println(string(result))
}
