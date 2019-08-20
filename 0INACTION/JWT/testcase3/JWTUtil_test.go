package testcase

import (
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	var (
		token string
		err   error
	)

	if token, err = GenerateToken(123456); err != nil {
		t.Fatalf("GenerateToken错误，错误信息：%s", err.Error())
	}
	t.Logf("token 值：%s", token)
}

func TestParseToken(t *testing.T) {
	var (
		err    error
		claims *CustomClaims
	)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjM0NTY3LCJleHAiOjE1NjI4MTAyNTIsImlzcyI6ImR1c3QifQ.OxY2F0mOX8Y8XhMPHcgxmXENyHAwg_i9eCWqokIw0QE"
	jwt := NewJWT()
	if claims, err = jwt.ParseToken(token); err != nil {
		t.Fatalf("ParseToken错误，错误信息：%s", err.Error())
	}
	t.Logf("claims中ID 值：%d", claims.ID)
}

func TestRefreshToken(t *testing.T) {
	var (
		err      error
		newToken string
	)
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIzNDU2LCJleHAiOjE1NjI3NTA3NzcsImlzcyI6ImR1c3QiLCJuYmYiOjE1NjI3NDk3MTd9.W-qTms5UeuBRtF6VfksaZeZANZWfP5NdFHwYPGDMp98"
	jwt := NewJWT()
	createToken, _ := GenerateToken(234567)
	time.Sleep(62 * time.Second)
	if _, err = jwt.ParseToken(createToken); err != nil && err == TokenExpired {
		if newToken, err = jwt.RefreshToken(createToken); err == nil {
			t.Logf("newToken 值：%s", newToken)
		} else {
			t.Fatalf("RefreshToken错误，错误信息：%s", err.Error())
		}
	}
}
