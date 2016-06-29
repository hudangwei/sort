package sort

// 最小堆排序后是递减数组
// 最大堆排序后是递增数组
// children of some index i: i*2+1, i*2+2
// parent of some index i: (i-1) / 2

func MinHeapFixup(arr []int, index int) {
	var parent int = (index - 1) / 2
	var tmp = arr[index]
	for parent >= 0 && index != 0 {
		if arr[parent] <= tmp {
			break
		}
		arr[index] = arr[parent]
		index = parent
		parent = (index - 1) / 2
	}
	arr[index] = tmp
}

func MinHeapAddNumber(arr []int, index int, number int) {
	arr[index] = number
	MinHeapFixup(arr, index)
}

func MinHeapFixdown(arr []int, index, n int) {
	var children = 2*index + 1
	var tmp = arr[index]
	for children < n {
		if children+1 < n && arr[children+1] < arr[children] {
			children++
		}
		if arr[children] >= tmp {
			break
		}
		arr[index] = arr[children]
		index = children
		children = 2*index + 1
	}
	arr[index] = tmp
}

func MinHeapDeleteNumber(arr []int, n int) {
	arr[0], arr[n-1] = arr[n-1], arr[0]
	MinHeapFixdown(arr, 0, n-1)
}

//arr 待排数组
//n 数组长度
func BuildMinHeap(arr []int, n int) {
	var index = n - 1
	var parent = (index - 1) / 2
	for i := parent; i >= 0; i-- {
		MinHeapFixdown(arr, i, n)
	}
}

func MinHeap_sort(arr []int, n int) {
	BuildMinHeap(arr, n)
	for i := n - 1; i >= 1; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		MinHeapFixdown(arr, 0, i)
	}
}

//----------------------------------------------------------------------------

func MaxHeapFixdown(arr []int, index, n int) {
	var children = 2*index + 1
	var tmp = arr[index]
	for children < n {
		if children+1 < n && arr[children+1] > arr[children] {
			children++
		}
		if arr[children] <= tmp {
			break
		}
		arr[index] = arr[children]
		index = children
		children = 2*index + 1
	}
	arr[index] = tmp
}

func MaxHeapDeleteNumber(arr []int, n int) {
	arr[0], arr[n-1] = arr[n-1], arr[0]
	MaxHeapFixdown(arr, 0, n-1)
}

//arr 待排数组
//n 数组长度
func BuildMaxHeap(arr []int, n int) {
	var index = n - 1
	var parent = (index - 1) / 2
	for i := parent; i >= 0; i-- {
		MaxHeapFixdown(arr, i, n)
	}
}

func MaxHeap_sort(arr []int, n int) {
	BuildMaxHeap(arr, n)
	for i := n - 1; i >= 1; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		MaxHeapFixdown(arr, 0, i)
	}
}
