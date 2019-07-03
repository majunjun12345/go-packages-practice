package api

import (
	"net/http"
	"testGoScripts/webFrameWork/echoWeb/common"
	"testGoScripts/webFrameWork/echoWeb/core"
	"testGoScripts/webFrameWork/echoWeb/db"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	u, err := core.GetOneUser()
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	}
	result := common.Struct2Map(*u)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}

func GetUser(c echo.Context) error {

	id := c.Param("id")
	u, err := core.GetUserByID(id)

	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	}
	result := common.Struct2Map(*u)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}

// -------------------------- mgo
// ---------user
const (
	// Key (Should come from somewhere else).
	Key = "secret"
)

func (h *db.Handler) SignUp(c echo.Context) error {
	return nil
}
