package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

/*
$ curl -v -X GET \
  http://localhost:8001 \
  -H 'Authorization: Basic dXNlcjpwYXNz' \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 82cbf224-b79b-4b80-b17e-b87af6ef446f'

Note: Unnecessary use of -X or --request, GET is already inferred.
* Rebuilt URL to: http://localhost:8001/
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8001 (#0)

> GET / HTTP/1.1
> Host: localhost:8001
> User-Agent: curl/7.47.0
> Accept: \*\/\*
> Authorization: Basic dXNlcjpwYXNz
> Cache-Control: no-cache
> Content-Type: application/json
> Postman-Token: 82cbf224-b79b-4b80-b17e-b87af6ef446f

< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< X-Powered-By: go-json-rest
< Date: Thu, 13 Jun 2019 00:54:36 GMT
< Content-Length: 23
<
* Connection #0 to host localhost left intact
{"Body":"Hello World!"}
*/

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultCommonStack...)
	api.Use(&rest.AuthBasicMiddleware{
		Realm: "my realm",
		Authenticator: func(userId, password string) bool {
			if userId == "user" && password == "pass" {
				return true
			}
			return false
		},
	})
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	}))
	log.Fatal(http.ListenAndServe(":8001", api.MakeHandler()))
}
