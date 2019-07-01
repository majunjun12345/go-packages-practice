package api

import (
	"net/http"
	"testGoScript/webFrameWork/echoWeb/common"
	"testGoScript/webFrameWork/echoWeb/core"

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