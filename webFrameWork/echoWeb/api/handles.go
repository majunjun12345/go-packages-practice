package api

import (
	"fmt"
	"log"
	"net/http"
	"testGoScripts/webFrameWork/echoWeb/common"
	"testGoScripts/webFrameWork/echoWeb/core"
	"testGoScripts/webFrameWork/echoWeb/models"
	"time"

	"github.com/dgrijalva/jwt-go"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

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

type (
	Handler struct {
		DB *mgo.Session
	}
)

var DBH *Handler

func InitMgoDB() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	if err := session.Copy().DB("twitter").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	DBH = &Handler{
		DB: session,
	}
}

func (h *Handler) SignUp(c echo.Context) error {
	u := &models.User{
		ID: bson.NewObjectId(),
	}
	if err := c.Bind(&u); err != nil {
		fmt.Println("11111")
		return err
	}
	fmt.Println("22222")
	// validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invald email or password"}
	}

	// save user
	db := h.DB.Clone()
	defer db.Close()
	if err := db.DB("twitter").C("users").Insert(u); err != nil { // 第一次看到这种方式的 save
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) Login(c echo.Context) error {
	var err error
	u := &models.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}

	db := h.DB.Clone()
	defer db.Close()
	if err := db.DB("twitter").C("users").Find(bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password!"}
		}
		return err
	}

	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjI3NTMzNDMsImlkIjoiNWQyMWMzYzA2ZjhmNTdlZGUzOGViYzRmIn0._S7XqLhJg3MF04dhcfXnl_P96Nk4B5XMg7MArFmY1Do
	u.Token, err = token.SignedString([]byte(Key))

	if err != nil {
		return err
	}
	u.Password = ""

	return c.JSON(http.StatusOK, u)
}

func (h *Handler) Follow(c echo.Context) error {
	fmt.Println("======")
	userID := userIDFromToken(c)
	id := c.Param("id")

	db := h.DB.Clone()
	defer db.Close()
	if err := db.DB("twitter").C("users").
		UpdateId(bson.ObjectIdHex(id), bson.M{"$addToSet": bson.M{"followers": userID}}); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
	}
	return nil
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	fmt.Println(c.Get("Authorization"))
	fmt.Println("user:", user)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
