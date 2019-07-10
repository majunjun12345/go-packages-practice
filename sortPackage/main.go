package main

import (
	"fmt"
	"sort"
)

func main() {
	// base()

	// structSort()

	structSort2()
}

// 基本类型的排序
func base() {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	// 升序
	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)
	fmt.Println(intList)
	fmt.Println(float8List)
	fmt.Println(stringList)

	// 降序
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Stable(sort.Reverse(sort.Float64Slice(float8List))) // stable：排序前后相同的元素位置不变
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))
	fmt.Println(intList)
	fmt.Println(float8List)
	fmt.Println(stringList)
}

// 结构体排序
type Person struct {
	Name string // 姓名
	Age  int    // 年纪
}
type PersonSlice []Person

func (a PersonSlice) Len() int           { return len(a) }
func (a PersonSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PersonSlice) Less(i, j int) bool { return a[i].Age < a[j].Age } // > 是降序

func structSort() {
	people := []Person{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}

	// 升序排列
	sort.Sort(PersonSlice(people))
	fmt.Println(people)

	// 降序排列
	sort.Sort(sort.Reverse(PersonSlice(people)))
	fmt.Println(people)
}

// 结构体直接排序
func structSort2() {
	data := []Person{
		{"Alice", 20},
		{"Bob", 15},
		{"Jane", 30},
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].Age > data[j].Age
	})
	fmt.Println(data)
}
