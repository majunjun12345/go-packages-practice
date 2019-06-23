package models

type User struct {
	Id       int    `json:id redis:"id"`
	UserName string `json:username redis:"username"`
	Email    string `json:email redis:"email"`
}

type Comment struct {
	Id          int    `json:"id" redis:"id" `
	User        User   `json:"user" redis:"user"`
	Text        string `json:"text" redis:"text"`
	CreatedTime string `json:"createdtime" redis:"createdtime"`
}

type Post struct {
	Id          int     `json:"id" redis:"id"`
	User        User    `json:"user" redis:"user"`
	Topic       string  `json:"topic" redis:"topic"`
	Text        string  `json:"text" redis:"text"`
	Comment     Comment `json:"comment" redis:"comment"`
	CreatedTime string  `json:"createdtime" redis:"createdtime"`
}
