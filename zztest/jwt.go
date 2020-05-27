package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const (
	TOKEN_DEVICE_ID = "orgi"
	TOKEN_THING_ID  = "thid"
	USER_ID         = "owur"
)

var publicKey = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDu0b92489luMOMDNSykdmco+SB
9kDHnCAUBw8h5Q0MoX3Wzs4yBFTL8n6Nfdd2gC40bxM7rVkYkAB2nowAYZ85C/F0
gpoldlnVs2HyG1vp0BX+ULuJQk756vpLfRlT5Kq0+tNsB+OKQuJ6lOyRKjNq9Q0+
qDPJOX/82N5QdT2b9QIDAQAB
-----END PUBLIC KEY-----
`

var tokenStr = `
eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY3IiOiIxIiwiYXVkIjoiaWFtIiwiYXpwIjoiaWFtIiwiZXhwIjoxNjE4MjEzODcxLCJpYXQiOjE1ODY2Nzc4NzEsImlzcyI6InN0cyIsImp0aSI6IjJEd1pReFV0ajRTMURsM3kxdG1IRTUiLCJuYmYiOjAsIm9yZ2kiOiJpb3RkLWI2ODk2ZGVhLTMzNzAtNGQ4Ni1hYjY1LTUyMTIzNzcwZDc4YiIsIm93dXIiOiJ1c3ItMnhoWEljN3AiLCJzdWIiOiJzdHMiLCJ0aGlkIjoiaW90dC05MFhTRUF1Y3htIiwidHlwIjoiSUQifQ.m_zHEVtGVZa5zjZIU8t3xC0QcIExeY5kaK-KhaTQGL7AQ0q3KtFMbPaOuFDwm3YtYEYB3YSUTz44RK1qY0UHW8l8Ds6eGXn_jhP8Xo-8D7Cfo-y_NCYT5CfYk_v9-mkXzDSYspWNjDdVuBRIHikamIcBXmiM5F2fWbghNK2VTLY
`

func Parse(tokenS string) error {
	if token == "" {
		return errors.New("token is blank")
	}

	// parse token
	token := string([]byte(tokenS))
	re := regexp.MustCompile(`\r?\n`)
	token = re.ReplaceAllString(token, "")
	token = strings.TrimSpace(token)

	pb, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey)) //解析公钥
	if err != nil {
		fmt.Println("ParseRSAPublicKeyFromPEM:", err.Error())
		return err
	}

	// parse token
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return pb, nil
	})

	if err != nil {
		panic(err)
	}

	if payload, ok := jwtToken.Claims.(jwt.MapClaims); ok {
		if err := payload.Valid(); err != nil {
			return err
		}

		fmt.Println(payload)
		_, ok := payload[TOKEN_DEVICE_ID].(string)
		if !ok {
			return errors.New("device id type error")
		}
		_, ok = payload[TOKEN_THING_ID].(string)
		if !ok {
			return errors.New("device id type error")
		}
		_, ok = payload[USER_ID].(string)
		if !ok {
			return errors.New("device id type error")
		}
		return nil
	}
	return errors.New("token error")
}
