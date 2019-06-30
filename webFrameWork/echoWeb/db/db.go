package db

import (
	"database/sql"
	"log"
	"testGoScripts/webFrameWork/echoWeb/models"

	_ "github.com/lib/pq"
)

/*
	关于数据库初始化的问题:
		定义全局变量 DB (mongo 中定义 Session 和 Collection 的全局变量)
		在 db 包层面定义 Init 函数, 用于初始化数据库
		在 app 包层面定义 init 函数, 调用各个包的 Init 函数, 进行统一初始化
*/

var DB *sql.DB

func Init() {
	db, err := sql.Open("postgres", "user=mamengli password=123456 dbname=mamengli sslmode=verify-full")
	if err != nil {
		log.Fatal("数据库连接失败:" + err.Error())
	}
	DB = db
}

// 查找一条记录
func FindOne() (*models.Excuse, error) {
	var id string
	var quote string
	err := DB.QueryRow("select id, quote from excuses LIMIT 1").Scan(&id, &quote)

	if err != nil {
		return nil, err
	}

	e := &models.Excuse{
		Id:    id,
		Quote: quote,
	}

	return e, nil
}
