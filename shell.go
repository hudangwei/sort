package sort

import (
	"math"
)

// 根据增量进行插入排序
// arr待排序数组
// n数组长度
// d增量
func shell_insertsort(arr []int, n int, d int) {
	for p := d; p < n; p++ {
		tmp := arr[p]
		var i int
		for i = p; i >= d && arr[i-d] > tmp; i -= d {
			arr[i] = arr[i-d]
		}
		arr[i] = tmp
	}
}

func Shell_sort1(arr []int, n int) {
	for d := n / 2; d > 0; d /= 2 {
		shell_insertsort(arr, n, d)
	}
}

//计算Hibbard增量
func dkHibbard(k float64) int {
	return int(math.Pow(2, k) - 1)
}

func Shell_sort2(arr []int, n int) {
	gaps := []int{7, 4, 1}
	for _, d := range gaps {
		shell_insertsort(arr, n, d)
	}
}
