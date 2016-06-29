package sort

import (
	"math/rand"
)

func initList() ([]int, []int) {
	unsorted := []int{11, 7, 33, 9, 12, 54, 3, 27, 41, 99, 2, 87}
	sorted := []int{2, 3, 7, 9, 11, 12, 27, 33, 41, 54, 87, 99}
	return unsorted, sorted
}

func makeslice(n int) []int {
	return rand.Perm(n)
}
