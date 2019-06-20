package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
	go 中大写开头的字段,在 mgo 中都变成了小写, 除非有标签
	分页:https://segmentfault.com/a/1190000019437349?utm_source=tag-newest
*/

var session mgo.Session
var db *mgo.Database
var c *mgo.Collection

func init() {
	session, err := mgo.Dial("localhost:27017")
	// mgo.DialWithTimeout
	// mgo.DialWithInfo() 可以设置用户名 密码 连接池(默认是 4096) 超时 数据库等
	db = session.DB("testGoMgo")
	/*
		结论:
			strong 可靠性最强, eventual 对读写分离性能的提升最强

		Strong(默认)
			session 的读写一直向主服务器发起并使用一个唯一的连接，因此所有的读写操作完全的一致。
		Monotonic
			session 的读操作开始是向其他服务器发起（且通过一个唯一的连接），只要出现了一次写操作，session 的连接就会切换至主服务器。
			由此可见此模式下，能够分散一些读操作到其他服务器，但是读操作不一定能够获得最新的数据。
		Eventual
			session 的读操作会向任意的其他服务器发起，多次读操作并不一定使用相同的连接，也就是读操作不一定有序。
			session 的写操作总是向主服务器发起，但是可能使用不同的连接，也就是写操作也不一定有序。
	*/
	// session.SetMode(mgo.Monotonic, true) // 设置 mongo 的读写模式
	// defer session.Close()   // 不能这样，执行 init 后就关闭了
	CheckErr(err)
	c = db.C("people")
}

func main() {
	// mgoInsert()

	// mgoSearch()

	mgoPage()

	// mgoUpdate()

	// mgoDel()
}

type Person struct {
	ID    bson.ObjectId `bson:"_id"`
	Name  string        `bson:"name"`
	Phone string        `bson:"phone"`
	Age   int           `bson:"age"`
}

func mgoInsert() {
	err := c.Insert(&Person{ID: bson.NewObjectId(), Name: "menglima", Phone: "189", Age: 99}, &Person{ID: bson.NewObjectId(), Name: "mamengli", Phone: "176", Age: 99})
	CheckErr(err)
}

func mgoSearch() {

	// one
	result := Person{}
	// err = c.Find(bson.M{"name": "menglima"}).One(&result)
	// err = c.FindId(bson.M{"_id": 12}).One(&result)
	err := c.Find(bson.M{"age": bson.M{"$lt": 20}}).One(&result) // lt:<, gt:>, 加上 e 就有了 =
	CheckErr(err)
	fmt.Printf("%+v\n", result)

	// all
	// res := []Person{}
	// err = c.Find(bson.M{"name": bson.M{"$regex": "A", "$options": "$i"}}).All(&res) // $regex 正则匹配, "$options": "$i" 忽略大小写, name 的 n 必须是小写
	// CheckErr(err)
	// fmt.Printf("%+v\n", res)
}

// -----------------分页
func mgoPage() {

	/*
		1. 适合数据量不大的情况，需要排序
	*/
	res := []Person{}
	/*
		selector：1 表示需要该字段，默认是 0
		sort：前面加上 - 表示倒叙，默认升序
	*/
	c.Find(bson.M{"name": "mamengli"}).Select(bson.M{"age": 1, "name": 1}).Sort("-age").Skip(0).Limit(150).All(&res)
	fmt.Printf("%+v", res)

	/*
		2. 数据量比较大，不需要排序，将上面的 sort 去掉就行
	*/

	/*
		3. 数据量大，需要排序，取得返回值后再进行排序
		pipeline 就是 mongo 里面的聚合
	*/
	// pipeM := []bson.M{
	// 	{"$match": bson.M{"name": "menglima"}},
	// 	{"$skip": 0},
	// 	{"$limit": 10},
	// 	{"$sort": bson.M{"age": 1}}, // 1 生序，-1 倒叙
	// }
	// pipe := c.Pipe(pipeM)
	// err := pipe.All(&res)
	// CheckErr(err)
	// fmt.Println(res)
}

func mgoUpdate() {
	p := Person{bson.NewObjectId(), "zhangsan", "111", 15}
	err := c.Update(bson.M{"name": "menglima"}, p)
	CheckErr(err)

	// 批量更新和只更新一条格式不一样
	updater := bson.M{"$set": bson.M{"age": 22}}
	info, err := c.UpdateAll(bson.M{"age": 22}, updater)
	CheckErr(err)
	fmt.Println(info.Updated)
}

func mgoDel() {
	err := c.Remove(bson.M{"age": 15})
	CheckErr(err)

	info, err := c.RemoveAll(bson.M{"age": 21})
	CheckErr(err)
	fmt.Println(info.Removed)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
