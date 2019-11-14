package testPackage

func Add(a, b int) int {
	return a + b
}

// bubble
func BubbleSort(nums []int) []int {
	l := len(nums)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

//选择排序
func SelectSort(a []int) []int {
	lenth := len(a)
	var minIndex int
	for i := 0; i < lenth; i++ {
		minIndex = i
		for j := i + 1; j < lenth; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
	return a
}

//插入排序
func InsertSort(a []int) []int {
	lenth := len(a)
	for i := 1; i < lenth; i++ {
		index := i - 1
		number := a[i]
		for index >= 0 && number < a[index] {
			a[index+1], a[index] = a[index], a[index+1]
			index--
		}
	}
	return a
}
