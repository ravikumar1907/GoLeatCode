package main

import (
	"fmt"
)

func partition(a []int, low, high int) int {
	p := a[low]
	i := low
	j := high
	for i < j {
		for a[i] < p && i < high {
			i++
		}
		for a[j] > p && j > low {
			j--
		}
		a[i], a[j] = a[j], a[i]
	}
	return j
}
func QuickSort(a []int, low, high int) {
	if low < high {
		p := partition(a, low, high)
		QuickSort(a, low, p)
		QuickSort(a, p+1, high)
	}
	return
}
func KthSmallest(a []int, low, high, k int) {
	if (high-low+1) < k || k <= 0 {
		return
	}
	for low <= high {
		p := partition(a, low, high)
		if p == k-1 {
			fmt.Println(a[k-1])
			return
		} else if p > k-1 {
			high = p - 1
		} else {
			low = p + 1
		}
	}
}

/* main() {
	a := []int{5, 6, 1, 7, 2, 8, 9, 3, 10, 4}
	fmt.Println(a)
	QuickSort(a, 0, 9)
	fmt.Println(a)
	a = []int{5, 6, 1, 7, 2, 8, 9, 3, 10, 4}
	fmt.Println(a)
	KthSmallest(a, 0, 9, 1)
	KthSmallest(a, 0, 9, 2)
	KthSmallest(a, 0, 9, 3)
	KthSmallest(a, 0, 9, 4)
	KthSmallest(a, 0, 9, 5)
	KthSmallest(a, 0, 9, 6)
	KthSmallest(a, 0, 9, 7)
	KthSmallest(a, 0, 9, 8)
	KthSmallest(a, 0, 9, 9)
	KthSmallest(a, 0, 9, 10)
}
*/
