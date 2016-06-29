package sort

func medianOfThree(arr []int, left, right int) int {
	var center int = (left + right) / 2
	if arr[left] > arr[center] {
		arr[left], arr[center] = arr[center], arr[left]
	}
	if arr[left] > arr[right] {
		arr[left], arr[right] = arr[right], arr[left]
	}
	if arr[center] > arr[right] {
		arr[center], arr[right] = arr[right], arr[center]
	}
	// 将pivot藏在右边，只需考虑left+1...right-2
	arr[center], arr[right-1] = arr[right-1], arr[center]
	return arr[right-1]
}

func quick_sort(arr []int, l, r int) {
	if r > l {
		pivot := medianOfThree(arr, l, r)
		var i, j = l + 1, r - 1
		for i < j {
			for i < j && arr[i] < pivot { //从左到右找第一个大于pivot的数
				i++
			}
			if i < j {
				arr[j] = arr[i]
				j--
			}
			for i < j && arr[j] >= pivot { //从右向左找第一个小于pivot的数
				j--
			}
			if i < j {
				arr[i] = arr[j]
				i++
			}
		}
		arr[j] = pivot
		quick_sort(arr, l, i-1)
		quick_sort(arr, i+1, r)
	}
}

func Quick_sort(arr []int, n int) {
	quick_sort(arr, 0, n-1)
}
