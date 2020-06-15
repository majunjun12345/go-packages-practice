package demo

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 测试当 c 为nil时，先进行一轮的垃圾回收，解绑 finalizer，下一轮 gc 将会彻底回收

func TestDemon3(t *testing.T) {
	s := assert.New(t)

	c := NewCache()
	cnt := 0
	stopped := make(chan struct{})
	c.onStopped = func() {
		cnt++
		close(stopped)
	}

	s.Equal(0, cnt)

	c = nil
	t.Log(cnt)

	runtime.GC()

	select {
	case <-stopped:
	case <-time.After(10 * time.Second):
		t.Fail()
	}
	t.Log(cnt)

	s.Equal(1, cnt)
}
