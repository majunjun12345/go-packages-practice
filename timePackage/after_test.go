package main

import (
	"testing"
	"time"
)

func TestAfter(t *testing.T) {

	time.After(time.Second * 3)
}
