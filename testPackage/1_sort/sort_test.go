package testPackage

import (
	"fmt"
	"testing"
)

// 初始化工作
func TestMain(m *testing.M) {
	fmt.Println("test main first!")
	m.Run()
}

// 单元测试, 开头字母大写才会被执行, 如果要使用 sub test,可以将开头字母改为小写
func testAdd(t *testing.T) {
	t.SkipNow()
	sum := Add(1, 2)
	if sum != 3 {
		t.Fail()
	}
}

// 表组测试
func testAddGroup(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{[]int{1, 2}, 3},
		{[]int{2, 3}, 5},
		{[]int{3, 3}, 6},
	}

	for _, test := range tests {
		sum := Add(test.data[0], test.data[1])
		if sum != test.want {
			t.Fatal()
		}
	}
}

func TestBubble(t *testing.T) {
	data := []int{1, 5, 3, 6, 4}
	BubbleSort(data)
}

// subtests, 强制测试的顺序执行,针对前后依赖
func TestAll(t *testing.T) {
	t.Run("TestAdd", testAdd)
	t.Run("TestAddGroup", testAddGroup)
}

