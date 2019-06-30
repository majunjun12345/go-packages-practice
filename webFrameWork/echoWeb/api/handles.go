package api

import (
	"net/http"
	"testGoScripts/webFrameWork/echoWeb/common"
	"testGoScripts/webFrameWork/echoWeb/core"

	"github.com/labstack/echo"
)

func index(c echo.Context) error {
	u := core.GetOneUser()
	if u == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": "",
		})
	}
	result := common.Struct2Map(u)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}
