package common_test

import (
	"testGoScript/webFrameWork/echoWeb/common"
	"testGoScript/webFrameWork/echoWeb/models"
	"testing"
)

func TestStruct2Map(t *testing.T) {
	e := models.Excuse{
		Id:    "123",
		Quote: "hello world",
	}

	m := common.Struct2Map(e)

	t.Log(m)

}
