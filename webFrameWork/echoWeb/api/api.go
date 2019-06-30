package api

import (
	"echo/middleware"
	"net/http"
	"testGoScripts/webFrameWork/echoWeb/db"

	"github.com/labstack/echo"
)

func init() {
	db.Init()
}

func StartServer() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"ping": "pong"})
	})

	e.GET("/", index)

	e.Logger.Fatal(e.Start(":8081"))
}
