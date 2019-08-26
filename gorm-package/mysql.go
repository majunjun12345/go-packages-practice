package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func StartMysql() {
	var err error

	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.Fatalf("open mysql err:%v", err)
	}

	db.SingularTable(true) // 全局禁用表名复数
	// open 默认返回的就是连接池，以下设置连接池配置
	db.DB().SetMaxIdleConns(10) // 最大备用连接数
	db.DB().SetMaxOpenConns(10) // 最大连接数，设置成并发量即可

	if !db.HasTable(&MessageUser{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&MessageUser{})
		db.Create(&MessageUser{
			WechatId: "menglima",
			Nickname: "12345654321",
			Avatar:   "617344533a2a",
		})
	}
	if !db.HasTable(&Message{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Message{})
		db.Create(&Message{
			// MessageUserID: 1,
			SenderId:   "11111111",
			ReceiverId: "menglima",
			Content:    "aaaaaaaa",
			IsNew:      true,
		})
		db.Create(&Message{
			SenderId:   "11111111",
			ReceiverId: "menglima",
			Content:    "bbbbbbbb",
			IsNew:      true,
		})
		db.Create(&Message{
			SenderId:   "11111111",
			ReceiverId: "menglima",
			Content:    "cccccccc",
			IsNew:      true,
		})
		db.Create(&Message{
			SenderId:   "sssss",
			ReceiverId: "sssss",
			Content:    "sssss",
			IsNew:      true,
		})
		db.Create(&Message{
			SenderId:   "rds",
			ReceiverId: "87654",
			Content:    "cccccccc",
			IsNew:      true,
		})
	}
}

func closeDB() {
	defer db.Close()
}
