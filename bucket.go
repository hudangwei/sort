package sort

const BUCKET_COUNT = 10

// 桶排序
func Bucket_sort(arr []int,n int)  {
    buckets := make([][]int, BUCKET_COUNT)
    for i := 0; i < n; i++ {
        idx := arr[i]/BUCKET_COUNT
        buckets[idx] = append(buckets[idx],arr[i])
    }
    index := 0
    bucketslen := len(buckets)
    for i := 0; i < bucketslen; i++ {
        if len(buckets[i]) == 0 {
            continue
        }
        Quick_sort(buckets[i],len(buckets[i]))
        for _,v := range buckets[i] {
            arr[index] = v
            index++
        }
    }
}