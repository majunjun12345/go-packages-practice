package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test_go?charset=utf8")
	defer db.Close()
	CheckErr(err)
}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
