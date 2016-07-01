package sort

//对于排序的关键字范围，一定是0-99
const NUM_RANGE = 100

// 计数排序
func Counting_sort(arr []int,n int)  {
    bucket := make([]int, NUM_RANGE)
    for i := 0; i < n; i++ {
        bucket[arr[i]]++
    }
    index := 0
    bucketlen := len(bucket)
    for i := 0; i < bucketlen; i++ {
        if bucket[i] == 0 {
            continue
        }
        for j := 0; j < bucket[i]; j++ {
            arr[index] = i
            index++
        }
    }
}