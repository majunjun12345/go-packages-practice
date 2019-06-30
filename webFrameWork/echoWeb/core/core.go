package core

import (
	"testGoScripts/webFrameWork/echoWeb/db"
	"testGoScripts/webFrameWork/echoWeb/models"
)

func GetOneUser() *models.Excuse {
	u, err := db.FindOne()
	if err != nil {
		return nil
	}
	return u
}
