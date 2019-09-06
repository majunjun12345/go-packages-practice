package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// https://studygolang.com/articles/16667
// 通过主表查关联表
// db.Where("wechat_id=?", "menglima").First(&messageUser)
// related方法需要第二个参数外键名称，如果没有第二个参数，就需要在你需要设置外键的地方加上一个字段userid
// db.Model(&user).Related(&emails)

func main() {
	StartMysql()
	GetAll()
	// GetOne()
}

func GetAll() {
	var messageUsers []MessageUser
	var data []MessageUser
	err := db.Where("wechat_id=?", "menglima").Find(&messageUsers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err)
	}
	for _, goodUser := range messageUsers {
		err = db.Model(&goodUser).Association("Messages").Find(&goodUser.Messages).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			fmt.Println(err)
		}
		fmt.Println("====message:", len(goodUser.Messages))
		data = append(data, goodUser)
	}
	fmt.Println("=======data:", len(data[0].Messages))
}

// 查找某个主表及对应的从表
func GetOne() {
	var messageUser MessageUser
	db.Where("wechat_id=?", "menglima").First(&messageUser)
	fmt.Println("======:", messageUser)
	err := db.Model(&messageUser).Association("Messages").Find(&messageUser.Messages).Error
	// err := db.Model(&messageUser).Related(&messageUser.Messages, "WechatId").Find(&messageUser.Messages)
	if err != nil {
		fmt.Println("=====:", err.Error)
	}
	fmt.Println(messageUser)
}
