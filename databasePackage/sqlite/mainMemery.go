package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/*
	事务:
	tx, err := db.Begin()
	tx.Prepare
	stmt.Exec
		defer stmt.Close()
	tx.Commit()
*/

type BCCode struct {
	b_code     string
	c_code     string
	code_type  int
	is_integer int
}

func main() {
	db, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared&loc=auto")
	defer db.Close()
	CheckErr1(err)

	fmt.Println("SQLite start")

	//创建表//delete from BC;，SQLite字段类型比较少，bool型可以用INTEGER，字符串用TEXT
	sqlStmt := `create table BC (
		b_code text not null primary key, 
		c_code text not null, 
		code_type INTEGER, 
		is_new INTEGER
		);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Println("create table error->%q: %s\n", err, sqlStmt)
		return
	}

	//创建索引，有索引和没索引性能差别巨大，根本就不是一个量级，有兴趣的可以去掉试试
	_, err = db.Exec("CREATE INDEX inx_c_code ON BC(c_code);")
	if err != nil {
		fmt.Println("create index error->%q: %s\n", err, sqlStmt)
		return
	}

	// 插入 10 万条数据
	start := time.Now().Unix()
	tx, err := db.Begin()
	CheckErr1(err)
	stmt, err := tx.Prepare("insert into BC(b_code, c_code, code_type, is_new) values(?,?,?,?)")
	CheckErr1(err)
	defer stmt.Close()
	var n int = 1000 * 1000
	for i := 0; i < n; i++ {
		_, err := stmt.Exec(fmt.Sprintf("B%024d", i), fmt.Sprintf("B%024d", i), 0, 1)
		CheckErr1(err)
	}
	tx.Commit()
	end := time.Now().Unix()

	// 随机检索 10 万次
	var count int = 0
	stmt, err = db.Prepare("select b_code, c_code, code_type, is_new from BC where c_code=?")
	defer stmt.Close()
	CheckErr1(err)
	bc := new(BCCode)
	for i := 0; i < n; i++ {
		err := stmt.QueryRow(fmt.Sprintf("B%024d", i)).Scan(&bc.b_code, &bc.c_code, &bc.code_type, &bc.is_integer)
		CheckErr1(err)
		count++
	}
	queryEnd := time.Now().Unix()
	fmt.Println("insert into 10万次,cost:", float64(end)-float64(start))
	fmt.Println("query 10万次,cost:", float64(queryEnd)-float64(end))
	fmt.Println(count)
}

func CheckErr1(err error) {
	if err != nil {
		panic(err)
	}
}
