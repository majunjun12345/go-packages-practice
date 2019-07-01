package core

import (
	"testGoScript/webFrameWork/echoWeb/db"
	"testGoScript/webFrameWork/echoWeb/models"
)

func GetOneUser() (*models.Excuse, error) {
	u, err := db.FindOne()
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByID(id string) (*models.Excuse, error) {
	u, err := db.FindOne()
	if err != nil {
		return nil, err
	}
	return u, nil
}
