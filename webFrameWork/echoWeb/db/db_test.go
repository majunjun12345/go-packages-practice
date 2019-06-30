package db_test

import (
	"testGoScript/webFrameWork/echoWeb/db"
	"testing"
)

func init() {
	db.Init()
}
func TestFindOne(t *testing.T) {
	u, err := db.FindOne()
	if err != nil {
		t.Fail()
	}
	t.Log(u)
}
