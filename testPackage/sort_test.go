package testPackage

import (
	"testing"
)

// 单元测试
func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if sum != 3 {
		t.Fail()
	}
}

// 表组测试
func TestAddGroup(t *testing.T) {
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

// 基准测试
func BenchmarkBubble(b *testing.B) {
	b.N = 2000000
	for i := 0; i < b.N; i++ {
		data := []int{1, 5, 3, 6, 4}
		BubbleSort(data)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	b.N = 2000000
	for i := 0; i < b.N; i++ {
		data := []int{1, 5, 3, 6, 4}
		SelectSort(data)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	b.N = 2000000
	for i := 0; i < b.N; i++ {
		data := []int{1, 5, 3, 6, 4}
		InsertSort(data)
	}
}
