package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
	go 和 mysql 中的 事件类型和bool 类型可以相互转换
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
		rows.Scan(&user.Id, &user.Username, &t, &n) // 和 query 一一对应,最好不要用 *
		fmt.Println(n)
	}
}

func GetTime() string {
	const shortForm = "2006-01-02 15:04:05"
	t := time.Now()
	return t.Format(shortForm)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
