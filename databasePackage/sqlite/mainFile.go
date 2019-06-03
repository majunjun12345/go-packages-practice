package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/*
	SQLite是个小型的数据库，很简洁，既支持文件也支持内存，比较适合小型的独立项目

	Exec 可用来执行 创建表、创建索引 等不需要传参的 sql
	需要传参的可以用 prepare + Exec 或 相关接受参数的 api

	DATE 字段和 time.Time 类型相呼应
*/
func mainFile() {
	db, err := sql.Open("sqlite3", "./foo.db")
	defer db.Close()
	CheckErr(err)

	sql_table := `
	CREATE TABLE IF NOT EXISTS userinfo(
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(64) NULL,
		departname VARCHAR(64) NULL,
		created DATE NULL
	)
	`
	_, err = db.Exec(sql_table)
	CheckErr(err)

	// insert
	stmt, err := db.Prepare("insert into userinfo(username, departname, created) values (?,?,?)")
	CheckErr(err)
	res, err := stmt.Exec("wangshubo", "国务院", "2017-04-21")
	CheckErr(err)
	id, _ := res.LastInsertId()
	rows, _ := res.RowsAffected()
	fmt.Printf("id:%d, rows:%d\n", id, rows)

	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	CheckErr(err)
	res, err = stmt.Exec("wangshubo_new", id)
	CheckErr(err)
	rows, _ = res.RowsAffected()
	fmt.Printf("rows:%d\n", rows)

	// query
	rowss, err := db.Query("select * from userinfo")
	CheckErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rowss.Next() {
		err = rowss.Scan(&uid, &username, &department, &created)
		CheckErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	rowss.Close() // 一定要 close

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	res, _ = stmt.Exec(10)
	affect, _ := res.RowsAffected()
	fmt.Println(affect)

	// 查询一条
	r := db.QueryRow("select * from userinfo where uid=?", 3)
	CheckErr(err)
	fmt.Println(r.Scan(&uid, &username, &department, &created))
	fmt.Println(uid, username, department, created)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
