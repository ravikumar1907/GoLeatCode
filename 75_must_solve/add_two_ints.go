package main

import (
	"unsafe"
)

func getSum(a, b int) int {
	for b != 0 {
		sum := a ^ b     // sum
		b = (a & b) << 1 // take carry and add it sum
		a = sum
	}
	return a
}

// max of 2 ints without using comparion operator
func maxOfTwoNumbers(a, b int) int {
	c := a - b
	size := unsafe.Sizeof(b) - 1
	k := (c << size) & 1
	return a - k*c
}

func topKFrequent(nums []int, k int) []int {
	var result [][]int
	hMap := make(map[int]int)
	for _, v := range nums {
		hMap[v] += 1
	}
	result = make([][]int, len(nums)+1)
	for k, v := range hMap {
		result[v] = append(result[v], []int{k}...)
	}
	tmp := make([]int, k)
	j := 0
	for i := len(nums) + 1; i > 0; i-- {
		if len(result[i]) != 0 {
			tmp = append(tmp, result[i]...)
			j += len(result[i])
			if j == k {
				break
			}
		}
	}
	return tmp
}

/*
func main() {
	fmt.Println(maxOfTwoNumbers(3, 5))
}
*/
