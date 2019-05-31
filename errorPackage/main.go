package main

import (
	"fmt"
	"testGoScripts/errorPackage/errorx"
	_ "testGoScripts/errorPackage/errorx"
)

type HttpError struct {
	Code    int
	Desc    string
	Message string
}

func NewHttpError(code int, desc, message string) *HttpError {
	return &HttpError{
		Code:    code,
		Desc:    desc,
		Message: message,
	}
}

func (he *HttpError) Error() string {
	return fmt.Sprintf("errcode:%s, errdesc:%s, detail message:%s", he.Code, he.Desc, he.Message)
}

func main() {
	a := 150
	newErr := errorx.IllegalState.New("unfortunate", "majun", "mamengli", a)
	fmt.Printf("%+v", errorx.Decorate(newErr, "this could be so much better"))
}
