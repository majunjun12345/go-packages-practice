package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
	插入:
		时间类型:
			go 先获取格式化的时间字符串, mysql 会根据事先定义的字段类型进行转换
		bool 类型:
			go 可以传递自己定义的 bool 类型, mysql 的 bool 类型会将其转换为 0(false) 和 1(true)
		id 自增:
			定义不赋值即可
	查询:
		从数据库中获取到的时间映射到 go 中是字符串, 所以 go 的接收类型必须是字符串
		数据库中获取到的 bool 类型映射到 go 中的是 int,所以 go 的接收类型必须是 int
		由于这两个特殊字段的关系,不太好直接映射到 go 中的 struct
*/

/*
	DROP TABLE IF EXISTS `userinfo`;
	CREATE TABLE `userinfo` (
		`id` INT(10) NOT NULL AUTO_INCREMENT,
		`username` VARCHAR(64) NULL DEFAULT NULL,
		`created` DATETIME NULL DEFAULT NULL,
		`married` BOOLEAN NOT NULL DEFAULT TRUE,
		PRIMARY KEY (`id`)
	);
*/

type UserInfo struct {
	Id          int       `db:"id"`
	Username    string    `db:"username"`
	CreatedTime time.Time `db:"created"`
	Married     bool      `db:"married"`
}

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/goTest?charset=utf8")
	// defer db.Close()
	CheckErr(err)
	DB = db
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数
}

func main() {
	// insert()

	// query()
}

// --------------------------insert
// 两种插入方式
func insert() {
	stmt, err := DB.Prepare("INSERT userinfo set username=?, created=?, married=?")
	CheckErr(err)
	_, err = stmt.Exec("memglima", GetTime(), false)
	CheckErr(err)

	DB.Exec("INSERT userinfo set username=?, created=?, married=?", "memglima", GetTime(), false)
}

// --------------------------query
func query() {
	/*
		返回的 row 必须使用 scan，不然会导致连接无法关闭，一直处于连接状态
		query 语句会建立一次连接，使用 scan 后会断开连接   https://www.jianshu.com/p/06f26f879d61
		* 也可以替换为具体的查询字段
	*/
	rows, err := DB.Query("select * from userinfo") // QueryRow 查询一行
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	if err != nil {
		panic(err)
	}
	var t string
	var n int
	for rows.Next() {
		user := UserInfo{}
		rows.Scan(&user.Id, &user.Username, &t, &n)
		fmt.Println(n)
	}
}

func GetTime() string {
	const shortForm = "2006-01-02 15:04:05"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	fmt.Println(t)
	return str
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
