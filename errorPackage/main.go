package main

import (
	"database/sql"
	"errors"
	"fmt"

	pkgErrors "github.com/pkg/errors"
)

var (
	ErrTestFailed = errors.New("test failed")
)

func main() {
	// A()
	AA()
}

// ----------------------------------------------------------------
func A() {
	err := B()

	// %v 只显示出文本信息
	fmt.Println(fmt.Sprintf("%v", pkgErrors.WithStack(err)))
	// %+v 能显示错误的堆栈信息
	fmt.Println(fmt.Sprintf("%+v", pkgErrors.WithStack(err)))
}

func B() error {
	return fmt.Errorf("this is %s error", "B")
}

func C() error {
	return fmt.Errorf("this is %s error", "C")
}

// ----------------------------------------------------------------

// ----------------------------------------------------------------

func AA() {
	err := Call()
	// Cause 返回最原始的 error
	if pkgErrors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found:%v\n", err)
	}
}

func Call() error {
	return pkgErrors.WithMessage(GetSql(), "bar failed")
}

func GetSql() error {
	return pkgErrors.Wrap(sql.ErrNoRows, "GetSql failed")
}

// ----------------------------------------------------------------
// grpc codes
func AAA() {
	// codes.Internal
	// status.
	
}
