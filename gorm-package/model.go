package main

import "github.com/jinzhu/gorm"

type MessageUser struct {
	gorm.Model

	Messages []Message `gorm:"FOREIGNKEY:ReceiverId;ASSOCIATION_FOREIGNKEY:WechatId"`
	WechatId string    `json:"wechat_id"`
	Nickname string    `json:"nickname"`
	Avatar   string    `json:"avatar"`
}

type Message struct {
	gorm.Model

	SenderId   string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
	Content    string `json:"message_content"`
	IsNew      bool   `json:"is_new"`
}
