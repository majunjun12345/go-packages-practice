package main

import "github.com/jinzhu/gorm"

// type MessageUser struct {
// 	gorm.Model

// 	Messages []Message `gorm:"FOREIGNKEY:ReceiverId;ASSOCIATION_FOREIGNKEY:WechatId"`
// 	WechatId string    `json:"wechat_id"`
// 	Nickname string    `json:"nickname"`
// 	Avatar   string    `json:"avatar"`
// }

// type Message struct {
// 	gorm.Model

// 	SenderId   string `json:"sender_id"`
// 	ReceiverId string `json:"receiver_id"`
// 	Content    string `json:"message_content"`
// 	IsNew      bool   `json:"is_new"`
// }
// 收件人的表
type MessageUser struct {
	gorm.Model

	Messages []Message `gorm:"FOREIGNKEY:ReceiverId;ASSOCIATION_FOREIGNKEY:WechatId"`
	WechatId string    `gorm:"type:VARCHAR(50);default:'';not null" json:"wechat_id"` // ReceiverId
	Nickname string    `gorm:"type:VARCHAR(20);default:'';column:昵称;not null" json:"nickname"`
	Avatar   string    `gorm:"type:VARCHAR(100);default:'';not null" json:"avatar"`
}

// 收件人信息表
type Message struct {
	gorm.Model

	SenderId   string `gorm:"type:VARCHAR(50);default:'';not null" json:"sender_id"`
	ReceiverId string `gorm:"type:VARCHAR(50);default:'';not null" json:"receiver_id"`
	Content    string `gorm:"type:TEXT" json:"message_content"`
	IsNew      int    `gorm:"type:int(1);default:1;not null" json:"is_new"` // 1 表示 new，2 表示已读
}
