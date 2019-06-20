package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session mgo.Session
var db *mgo.Database
var c *mgo.Collection

func init() {
	session, err := mgo.Dial("localhost:27017")
	db = session.DB("testGoMgo")
	defer session.Clone()
	CheckErr(err)
	c = db.C("people")
}

func main() {
	mgoSearch()
}

type Person struct {
	Name  string
	Phone string
	Age   int
}

func mgoInsert() {
	err := c.Insert(&Person{"menglima", "189", 21}, &Person{"mamengli", "176", 19})
	CheckErr(err)
}

func mgoSearch() {

	result := Person{}
	// err = c.Find(bson.M{"name": "menglima"}).One(&result)
	// err = c.FindId(bson.M{"_id": 12}).One(&result)
	res := []Person{}
	err := c.Find(bson.M{"name": bson.M{"$regex": "A", "$options": "$i"}}).All(&res) // $regex 正则匹配, "$options": "$i" 忽略大小写
	CheckErr(err)
	fmt.Printf("%+v\n", result)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
