package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "newtrekWang"
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	ID    string `json:"userId"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

//创建token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	res, err := token.SignedString(j.SigningKey)
	fmt.Println("err:", err)
	return res, err
}

//解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {

	// ParseWithClaims: 用于解析鉴权的声明，内部方法进行了解码和校验(sign)，最终返回 *jwt.Token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			log.Panicln("unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.SigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	// tokenClaims.Valid 用于验证基于时间的声明exp, iat, nbf
	// 如果在 token 中没有任何声明，仍被认为有效
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

//更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}

	return "", TokenInvalid
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	generateToken(w)
}

func generateToken(w http.ResponseWriter) {
	j := &JWT{[]byte("man")}
	claims := CustomClaims{
		"1", "Jaya", "123456", jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 3600),
			Issuer:    "man",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		io.WriteString(w, "it is wrong")
	}

	io.WriteString(w, token)
}

func Register() *httprouter.Router {
	router := httprouter.New()
	// router.POST("/reg", Reg)
	router.POST("/login", Login)
	return router
}

func main() {
	router := Register()
	http.ListenAndServe(":8005", router)
}
