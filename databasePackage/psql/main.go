package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "gotest"
)

var DB *sql.DB

type Teacher struct {
	ID   int
	Name string
	Age  int
}

func init() {
	// 如果没有密码则不写密码
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	CheckErr(err)
	DB = db
}

func main() {
	// Insert()
	Search()
}

func Insert() {
	defer DB.Close()

	sqlStatement := `
					INSERT INTO teacher (id, name, age)  
					VALUES ($1, $2, $3)  
					RETURNING id
					`
	stmt, err := DB.Prepare(sqlStatement)
	CheckErr(err)
	result, err := stmt.Exec(2, "menglima", 3)
	CheckErr(err)
	fmt.Println(result.RowsAffected())
}

func Search() {
	defer DB.Close()
	// 最好别使用 *，于下面的映射相对应
	sqlStatement := `SELECT id, name, age FROM teacher WHERE id=$1;`
	query := DB.QueryRow(sqlStatement, 1)
	var id int
	var name string
	var age int
	err := query.Scan(&id, &name, &age)
	CheckErr(err)
	fmt.Println(id, name, age)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
