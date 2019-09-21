package testPackage

import "testing"

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
