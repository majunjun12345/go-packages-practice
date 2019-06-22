package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/goTest?charset=utf8&parseTime=true")
	// defer db.Close()
	CheckErr(err)
	DB = db
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	v1 := r.Group("/api/v1/userinfo")
	{
		v1.GET("/user/:id", GetUser)
		v1.POST("/user", CreateUser)
		v1.GET("/users", FetchAllUsers)
		v1.PUT("/user/:id", UpdateUser)
		v1.DELETE("/user/:id", DeleteUser)
	}
	r.Run(":8081")
}

type User struct {
	ID       int       `db:"id"`
	UserName string    `db:"username"`
	Created  time.Time `db:"created"`
	Married  bool      `db:"married"`
}

func GetUser(c *gin.Context) {
	var user User
	var result gin.H
	id := c.Param("id") // 这里获取的 id 是 string 类型,但是数据库中的是 int 类型,貌似不用转换
	row := DB.QueryRow("select username, created, married from userinfo where id=?", id)
	err := row.Scan(&user.UserName, &user.Created, &user.Married)
	if err != nil {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": user,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

func CreateUser(c *gin.Context) {
	var result gin.H
	stmt, err := DB.Prepare("insert into userinfo (username, created, married) values (?, ?, ?)")
	CheckErr(err)
	_, err = stmt.Exec("zhangsan", GetTime(), false)
	CheckErr(err)
	result = gin.H{
		"result": "success",
	}
	c.JSON(http.StatusOK, result)

}

func FetchAllUsers(c *gin.Context) {
	var user User
	var users []User
	var result gin.H
	rows, err := DB.Query("select username, created, id, married from userinfo")
	if err != nil {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
		c.JSON(http.StatusOK, result)
	}

	// var create time.Time
	// var m int
	for rows.Next() {
		err := rows.Scan(&user.UserName, &user.Created, &user.ID, &user.Married) // 这里的字段必须和上面的一一对应, 如果 query *,则完全不知道顺序
		CheckErr(err)
		// if m == 1 {
		// 	user.Married = true
		// } else {
		// 	user.Married = false
		// }
		users = append(users, user)
	}

	result = gin.H{
		"result": users,
		"count":  len(users),
	}
	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}

func GetTime() string {
	const shortForm = "2006-01-02 15:04:05"
	t := time.Now()
	return t.Format(shortForm)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
