package main

import (
	"testing"

	"go.uber.org/goleak"
)

/*
	能在编译部署前识别 goroutine 泄漏的工具
*/

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestT(t *testing.T) {
	defer goleak.VerifyNone(t)
	main()
}
