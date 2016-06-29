package sort

func Insertion_sort(arr []int, n int) {
	for p := 1; p < n; p++ {
		tmp := arr[p]
		var i int
		for i = p; i >= 1 && arr[i-1] > tmp; i-- {
			arr[i] = arr[i-1]
		}
		arr[i] = tmp
	}
}
