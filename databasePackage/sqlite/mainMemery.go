package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared&loc=auto")
	defer db.Close()
	CheckErr(err)

	fmt.Println("SQLite start")

	//创建表//delete from BC;，SQLite字段类型比较少，bool型可以用INTEGER，字符串用TEXT
	sqlStmt := `create table BC (b_code text not null primary key, c_code text not null, code_type INTEGER, is_new INTEGER);`
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

}
