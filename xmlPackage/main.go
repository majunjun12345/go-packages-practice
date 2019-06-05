package main

import (
	"encoding/xml"
	"fmt"
)

type Address struct {
	City, State string
}

type Person struct {
	XMLName   xml.Name `xml:"person"`     // 相当于是定义 根元素
	Id        int      `xml:"id,attr"`    // attr 表示属性，上一个元素的，id 和 attr 之间不能有空格
	FirstName string   `xml:"name>first"` // name 下面有两个子元素
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height, omitempty"`
	Married   bool
	Address
	Comment string `xml:",comment"` // 加，和不加有区别
}

func main0() {
	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, " ", " ")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Println(string(output))
	fmt.Println("\n")
	/*
	 <person>
	  <attr xmlns="id,">13</attr>
	  <name>
	   <first>John</first>
	   <last>Doe</last>
	  </name>
	  <age>42</age>
	  <omitempty xmlns="height,">0</omitempty>
	  <Married>false</Married>
	  <City>Hanga Roa</City>
	  <State>Easter Island</State>
	  <!-- Need more details. -->
	 </person>
	*/

	// Unmarshal
	data := `
        <Person>
            <FullName>Grace R. Emlin</FullName>
            <Company>Example Inc.</Company>
            <Email where="home">
                <Addr>gre@example.com</Addr>
            </Email>
            <Email where='work'>
                <Addr>gre@work.com</Addr>
            </Email>
            <Group>
                <Value>Friends</Value>
                <Value>Squash</Value>
            </Group>
            <City>Hanga Roa</City>
            <State>Easter Island</State>
        </Person>
	`
	v1 := Result{
		Name:  "none",
		Phone: "none",
	}
	err = xml.Unmarshal([]byte(data), &v1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", v1)
	/*
		{XMLName:{Space: Local:Person}
		Name:Grace R. Emlin
		Phone:none
		Email:[{Where:home Addr:gre@example.com} {Where:work Addr:gre@work.com}]
		Groups:[]
		Address1:{City:Hanga Roa State:Easter Island}}
	*/
}

type Email struct {
	Where string `xml:"where,attr"`
	Addr  string
}
type Address1 struct {
	City, State string
}
type Result struct {
	XMLName xml.Name `xml:"Person"`
	Name    string   `xml:"FullName"`
	Phone   string
	Email   []Email
	Groups  []string `xml:"Group:Value"`
	Address1
}
