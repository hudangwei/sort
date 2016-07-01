package sort

func getDigit(num int,n int) int{
    return (num/n)%10
}

// 基数排序 LSD次位优先算法
func Radix_sort(arr []int,n int){
    buckets := make([][]int, BUCKET_COUNT)
    for i := 0; i < n; i++ {
        idx := getDigit(arr[i],1) //个位数
        buckets[idx] = append(buckets[idx],arr[i])
    }
    index := 0
    bucketslen := len(buckets)
    for i := 0; i < bucketslen; i++ {
        if len(buckets[i]) == 0 {
            continue
        }
        for _,v := range buckets[i] {
            arr[index] = v
            index++
        }
    }

    buckets2 := make([][]int, BUCKET_COUNT)
    for i := 0; i < n; i++ {
        idx := getDigit(arr[i],10) //十位数
        buckets2[idx] = append(buckets2[idx],arr[i])
    }
    index = 0
    buckets2len := len(buckets2)
    for i := 0; i < buckets2len; i++ {
        if len(buckets2[i]) == 0 {
            continue
        }
        for _,v := range buckets2[i] {
            arr[index] = v
            index++
        }
    }    
}