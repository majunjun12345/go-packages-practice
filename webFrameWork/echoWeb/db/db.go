package db

import (
	"database/sql"
	"echo"
	"fmt"
	"log"
	"testGoScripts/webFrameWork/echoWeb/db"
	"testGoScripts/webFrameWork/echoWeb/models"

	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
)

/*
	关于数据库初始化的问题:
		定义全局变量 DB (mongo 中定义 Session 和 Collection 的全局变量)
		在 db 包层面定义 Init 函数, 用于初始化数据库
		在 app 包层面定义 init 函数, 调用各个包的 Init 函数, 进行统一初始化
*/

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "test"
	password = "123456"
	// sslmode = "verify-full"
)

var DB *sql.DB

func Init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	er := db.Ping()
	if err != nil && er != nil { // 这里可能是 nil 但是数据库没有连接上, 应该 ping 一下
		log.Fatalf("数据库连接失败, err:, er:%s", err.Error(), er.Error())
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

func GetUserByID(id string) (*models.Excuse, error) {
	var quote string
	err := DB.QueryRow("select id, quote from excuses where id=?", id).Scan(&id, &quote)

	if err != nil {
		return nil, err
	}

	e := &models.Excuse{
		Id:    id,
		Quote: quote,
	}

	return e, nil
}

// -------------------------------------------- 使用 mgo

type (
	Handler struct {
		DB *mgo.Session
	}
)

var DBH *Handler

func InitMgoDB() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	if err := session.Copy().DB("twitter").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	DBH = &Handler{
		DB: session,
	}
}

func (h *db.Handler) SignUp(c echo.Context) error {
	return nil
}
