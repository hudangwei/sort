package sort

import (
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	unsorted, sorted := initList()
	Bubble_sort1(unsorted, len(unsorted))
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("bubble sort error")
		}
	}
}

func TestBubbleSort2(t *testing.T) {
	unsorted, sorted := initList()
	Bubble_sort2(unsorted, len(unsorted))
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("bubble sort error")
		}
	}
}

func TestBubbleSort3(t *testing.T) {
	unsorted, sorted := initList()
	Bubble_sort3(unsorted, len(unsorted))
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("bubble sort error")
		}
	}
}

func TestInsertionSort(t *testing.T) {
	unsorted, sorted := initList()
	Insertion_sort(unsorted, len(unsorted))
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("insertion sort error")
		}
	}
}

func TestMergeSort(t *testing.T) {
	unsorted, sorted := initList()
	Merge_sort(unsorted, len(unsorted))
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("merge sort error")
		}
	}
}

func TestQuickSort(t *testing.T) {
	unsorted, sorted := initList()
	Quick_sort(unsorted, len(unsorted))
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("quick sort error")
		}
	}
}

func TestSelectSort(t *testing.T) {
	unsorted, sorted := initList()
	Select_sort(unsorted, len(unsorted))
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("select sort error")
		}
	}
}

func TestShellSort1(t *testing.T) {
	unsorted, sorted := initList()
	Shell_sort1(unsorted, len(unsorted))
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("shell1 sort error")
		}
	}
}

func TestShellSort2(t *testing.T) {
	unsorted, sorted := initList()
	Shell_sort2(unsorted, len(unsorted))
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("shell2 sort error")
		}
	}
}

func TestMaxHeapSort(t *testing.T) {
	unsorted, sorted := initList()
	MaxHeap_sort(unsorted, len(unsorted))
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("minheap sort error")
		}
	}
}

func TestCountingSort(t *testing.T)  {
	unsorted, sorted := initList()
	Counting_sort(unsorted, len(unsorted))
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("counting sort error")
		}
	}
}

func TestBucketSort(t *testing.T)  {
	unsorted, sorted := initList()
	Bucket_sort(unsorted, len(unsorted))
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("bucket sort error")
		}
	}
}

func TestRadixSort(t *testing.T)  {
	unsorted, sorted := initList()
	Radix_sort(unsorted, len(unsorted))
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("radix sort error")
		}
	}
}