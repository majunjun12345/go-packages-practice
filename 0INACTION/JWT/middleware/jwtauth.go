package middleware

import (
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "github.com/pkg/errors"
)

type JWT struct {
    SigningKey []byte
}

var (
    TokenExpired     error  = errors.New("Token is expired")
    TokenNotValidYet error  = errors.New("Token not active yet")
    TokenMalformed   error  = errors.New("That's not even a token")
    TokenInvalid     error  = errors.New("Couldn't handle this token:")
    SignKey          string = "hlms_yeeuu"
)

// 自定义结构体参数
type CustomClaims struct {
    User      string `form:"phone"`
    Level     string `form:"type"`
    HotelId   string `form:"hotel_id"`
    HotelName string `form:"hotelName"`
    UserName  string `form:"username"`
    jwt.StandardClaims
}

// JWT验证
func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        token, err := c.Cookie("claims")
        if err != nil {
            c.Redirect(302, "/login")
            return
        }
        j := NewJWT()
        // token := t.(string)
        // claims, err := j.ParseToken(token)
        claims, err := j.ParseToken(token)
        if err != nil {
            if err == TokenExpired {
                if token, err = j.RefreshToken(token); err == nil {
                    c.JSON(http.StatusOK, gin.H{"error": 0, "message": "refresh token", "token": token})
                    return
                }
            }
            c.Redirect(302, "/login")
            return
        }
        c.Set("claims", claims)
        c.Next()
        return
    }
}

func NewJWT() *JWT {
    return &JWT{
        []byte(GetSignKey()),
    }
}
func GetSignKey() string {
    return SignKey
}
func SetSignKey(key string) string {
    SignKey = key
    return SignKey
}

// parse
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return j.SigningKey, nil
    })
    if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors&jwt.ValidationErrorMalformed != 0 {
                return nil, TokenMalformed
            } else if ve.Errors&jwt.ValidationErrorExpired != 0 {
                // Token is expired
                return nil, TokenExpired
            } else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
                return nil, TokenNotValidYet
            } else {
                return nil, TokenInvalid
            }
        }
    }
    if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, TokenInvalid
}
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

// create
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(j.SigningKey)
}

//------------------------------------------------------JWT调用(gin)------------------------------------------------------
// 生成
    jwt := middleware.NewJWT()
    token, err := jwt.CreateToken(claims)
    if err != nil {
        c.JSON(200, gin.H{"status": 206, "msg": err.Error()})
        return
    }
    c.SetCookie("claims", token, 80000, "/", "", false, true)
    c.JSON(200, gin.H{"status": 200, "msg": ""})


// 解析
//公共的获取cookie
func GetJWT(c *gin.Context) (customclaims *middleware.CustomClaims, err error) {
    claims, err := c.Cookie("claims")
    if err != nil {
        fmt.Println(errors.Wrap(err, "utils/GetJWT/获取失败"))
        return
    }
    jwt := middleware.NewJWT()
    customclaims, err = jwt.ParseToken(claims)
    if customclaims == nil {
        err = errors.New("JWT为nil")
    }
    return
}

// router 验证
    admin := app.Group("/admin", middleware.AdminJWTAuth())
    app.Use(middleware.JWTAuth())