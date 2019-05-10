package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
	该库还是执行原生的 sql 语句
*/

type AnnouncementTable struct { // 使用结构体解析数据库数据
	ID         int            `db:"id"`
	ImgUrl     string         `db:"imgUrl"`
	DetailUrl  sql.NullString `db:"detailUrl"`
	CreateDate string         `db:"createDate"`
	State      int            `db:"state"`
}

func main() {
	conn, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/goTest")
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	// 返回的 row 必须使用 scan，不然会导致连接无法关闭，一直处于连接状态
	// query 语句会建立一次连接，使用 scan 后会断开连接   https://www.jianshu.com/p/06f26f879d61

	rows, err := conn.Query("select * from announcement") // QueryRow 查询一行
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		announce := AnnouncementTable{}
		rows.Scan(&announce.ID, &announce.ImgUrl, &announce.DetailUrl, &announce.CreateDate, &announce.State)
		fmt.Println(announce)
	}
}
