package sort

// 冒泡排序1
// 待排序数组arr，数组长度n
func Bubble_sort1(arr []int, n int) {
	for j := n - 1; j >= 0; j-- {
		for i := 0; i < j; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

// 冒泡排序2
func Bubble_sort2(arr []int, n int) {
	for j := n - 1; j >= 0; j-- {
		flag := true
		for i := 0; i < j; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

// 冒泡排序3
func Bubble_sort3(arr []int, n int) {
	var k int
	flag := n - 1
	for flag > 0 {
		k = flag
		flag = 0
		for i := 0; i < k; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				flag = i
			}
		}
	}
}
