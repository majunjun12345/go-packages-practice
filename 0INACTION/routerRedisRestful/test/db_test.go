package test

import (
	"fmt"
	"testGoScript/0INACTION/routerRedisRestful/db"
	"testGoScript/0INACTION/routerRedisRestful/models"
	"testing"
)

func TestInsert(t *testing.T) {
	u := models.User{UserName: "menglima", Email: "15527254815@qq.com"}
	err := db.Insert(&u)
	if err != nil {
		fmt.Println(err)
		fmt.Println("end")
	}
}
