package api

import (
	"math/rand"
	"net/http"
	"testGoScript/webFrameWork/echoWeb/db"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	db.Init()
}

func StartServer() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"*"},
		AllowOrigins: []string{"http://foo.com", "http://test.com"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.Static("/", "static")

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"ping": "pong"})
	})

	e.GET("/", Index)
	e.GET("/id/:id", GetUser)

	// test cors
	/*
		curl -v -H 'Origin: http://foo.com' http://localhost:1323/api/users
		*   Trying ::1...
		* TCP_NODELAY set
		* Connected to localhost (::1) port 8081 (#0)
		> GET /api/users HTTP/1.1
		> Host: localhost:8081
		> User-Agent: curl/7.54.0
		> Accept: *-/*
		> Origin: http://foo.com
		>
		< HTTP/1.1 200 OK
		< Access-Control-Allow-Origin: http://foo.com
		< Content-Type: application/json; charset=UTF-8
		< Vary: Origin
		< Date: Mon, 01 Jul 2019 06:20:55 GMT
		< Content-Length: 22
		<
		["Joe","Veer","Zion"]
		* Connection #0 to host localhost left intact
	*/
	e.GET("/api/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []string{"Joe", "Veer", "Zion"})
	})

	// server pusher
	e.GET("/api/pusher", func(c echo.Context) error {
		pusher, ok := c.Response().Writer.(http.Pusher)

		if ok {
			if err := pusher.Push("/echo.jpeg", nil); err != nil {
				return err
			}

		}
		return c.File("static/index.html")
	})

	// jsonp 解决跨域问题
	e.GET("/jsonp", func(c echo.Context) error {
		callback := c.QueryParam("callback")
		var content struct {
			Response  string    `json:"response"`
			Timestamp time.Time `json:"timestamp"`
			Random    int       `json:"random"`
		}
		content.Response = "Sent via JSONP"
		content.Timestamp = time.Now().UTC()
		content.Random = rand.Intn(1000)
		return c.JSONP(http.StatusOK, callback, &content)
	})

	//使用 https and http/2.0
	/*
		curl -k -i https://localhost:8081/api/users 必带参数：-k

		HTTP/2 200
		access-control-allow-origin:
		content-type: application/json; charset=UTF-8
		vary: Origin
		content-length: 22
		date: Mon, 01 Jul 2019 09:37:53 GMT

		["Joe","Veer","Zion"]

	*/

	// e.Logger.Fatal(e.StartTLS(":8081", "pem/cert.pem", "pem/key.pem"))
	e.Logger.Fatal(e.Start(":8081"))

}
