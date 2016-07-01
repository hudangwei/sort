package sort

import (
	"testing"
)

func BenchmarkBubbleSort1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{5, 4, 8, 2, 1, 6, 7, 3, 9, 10}
		Bubble_sort1(arr, 10)
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{5, 4, 8, 2, 1, 6, 7, 3, 9, 10}
		Bubble_sort2(arr, 10)
	}
}

func BenchmarkBubbleSort3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{5, 4, 8, 2, 1, 6, 7, 3, 9, 10}
		Bubble_sort3(arr, 10)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{5, 4, 8, 2, 1, 6, 7, 3, 9, 10}
		Insertion_sort(arr, 10)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{5, 4, 8, 2, 1, 6, 7, 3, 9, 10}
		Merge_sort(arr, 10)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{5, 4, 8, 2, 1, 6, 7, 3, 9, 10}
		Quick_sort(arr, 10)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{5, 4, 8, 2, 1, 6, 7, 3, 9, 10}
		Select_sort(arr, 10)
	}
}

func BenchmarkShellSort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{5, 4, 8, 2, 1, 6, 7, 3, 9, 10}
		Shell_sort1(arr, 10)
	}
}

func BenchmarkShellSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{5, 4, 8, 2, 1, 6, 7, 3, 9, 10}
		Shell_sort2(arr, 10)
	}
}

func BenchmarkMinHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{5, 4, 8, 2, 1, 6, 7, 3, 9, 10}
		MinHeap_sort(arr, 10)
	}
}

func BenchmarkMaxHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{5, 4, 8, 2, 1, 6, 7, 3, 9, 10}
		MaxHeap_sort(arr, 10)
	}
}

func BenchmarkCountingSort(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		arr := []int{5, 4, 8, 2, 1, 6, 7, 3, 9, 10}
		Counting_sort(arr, 10)
	}
}

func BenchmarkBucketSort(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		arr := []int{5, 4, 8, 2, 1, 6, 7, 3, 9, 10}
		Bucket_sort(arr, 10)
	}
}

func BenchmarkRadixSort(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		arr := []int{15, 34, 68, 72, 11, 6, 97, 83, 29, 40}
		Radix_sort(arr, 10)
	}
}