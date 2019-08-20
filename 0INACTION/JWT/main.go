package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/urfave/negroni"
)

/*
	JWT 包含三部分：
	头：
		{
			"alg": "HS256",
			"typ": "JWT"
		}
		使用 base64 url 得出字符串
	有效载荷(json对象)：
		iss：发行人
		exp：到期时间
		sub：主题
		aud：用户
		nbf：在此之前不可用
		iat：发布时间
		jti：JWT ID用于标识该JWT
		自定义字段：
			"sub": "1234567890",
			"name": "chongchong",
			"admin": true
		使用 base64 url 得出字符串，未加密，勿放私密信息
	签名哈希：
		使用 secret 对以上两部分使用指定签名方法进行签名

	上述三个部分字符串组合，使用 . 连接
	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjA0OTI4MTgsImlhdCI6MTU2MDQ4OTIxOH0.L8FgH3846II40WQnsWfzN7Dbdhdb2K9rYpA9bPNzjb8

	用法：
	客户端：
		服务端先以 json 形式返回 token
		客户端下次请求将其放在请求头的 Authorization: Bearer 中能够实现跨域问题；
		Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjA0OTI4MTgsImlhdCI6MTU2MDQ4OTIxOH0.L8FgH3846II40WQnsWfzN7Dbdhdb2K9rYpA9bPNzjb8
	服务端：
		校验 token

	优势：
		jwt token 认证机制是对 session 认证机制的替代
		session 存储于服务端，用户过多会造成服务器压力；
		session 依赖 cookie，如果 cokkie 被截获，可能会造成 CSRF(跨站请求伪造)；
		不适合做分布式场景；

	当前脚本的校验直接通过 request.ParseFromRequest(),
	其余的脚本都是通过 jwt 包自带的认证方式：jwt.ParseWithClaims
*/

const (
	SecretKey = "welcome to menglima's blog"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Data string `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {

	response := Response{"Gained access to protected resource"}
	JsonResponse(response, w)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user UserCredentials

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}

	if strings.ToLower(user.Username) != "someone" {
		if user.Password != "p@ssword" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(w, "Invalid credentials")
			return
		}
	}

	// header + preload jwt.New() 使用默认 claims, NewWithClaims 使用自定义 claims
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	// header + preload + signedString
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		fatal(err)
	}

	response := Token{tokenString}
	JsonResponse(response, w)

}

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	fmt.Printf("token:%+v", token)
	/*
		{Raw:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjA0OTI4MTgsImlhdCI6MTU2MDQ4OTIxOH0.L8FgH3846II
		40WQnsWfzN7Dbdhdb2K9rYpA9bPNzjb8
		Method:0xc0000a82e0
		Header:map[alg:HS256 typ:JWT]    first segment
		Claims:map[exp:1.560492818e+09 iat:1.560489218e+09]   second segment
		Signature:L8FgH3846II40WQnsWfzN7Dbdhdb2K9rYpA9bPNzjb8   third segment
		Valid:true}
	*/
	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}

}

func JsonResponse(response interface{}, w http.ResponseWriter) {

	responseData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}

func StartServer() {

	http.HandleFunc("/login", LoginHandler)

	http.Handle("/resource", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
	))

	log.Println("Now listening...")
	http.ListenAndServe(":1234", nil)
}

func main() {
	StartServer()
}
