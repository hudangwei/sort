package sort

func Merge_sort(arr []int, n int) {
	tmp := make([]int, n)
	mergesort(arr, 0, n-1, tmp)
}

func mergesort(arr []int, l, r int, tmp []int) {
	if l < r {
		m := (l + r) / 2
		mergesort(arr, l, m, tmp)     //左边有序
		mergesort(arr, m+1, r, tmp)   //右边有序
		mergearray(arr, l, m, r, tmp) //将两个有序数组合并
	}
}

func mergearray(arr []int, l, mid, r int, tmp []int) {
	i, j := l, mid+1
	m, n := mid, r
	var k = 0
	for i <= m && j <= n {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			k++
			i++
		} else {
			tmp[k] = arr[j]
			k++
			j++
		}
	}
	for i <= m {
		tmp[k] = arr[i]
		k++
		i++
	}
	for j <= n {
		tmp[k] = arr[j]
		k++
		j++
	}
	for i = 0; i < k; i++ {
		arr[l+i] = tmp[i]
	}
}
