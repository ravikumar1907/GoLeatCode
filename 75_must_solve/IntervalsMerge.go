package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	var result [][]int
	n := len(intervals)
	if n < 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	result = append(result, intervals[0])
	for i := 1; i < n; i++ {
		a := result[len(result)-1]
		b := intervals[i]
		if b[0] > a[1] {
			result = append(result, b)
		} else {
			if a[1] < b[1] {
				a[1] = b[1]
			}
		}
	}
	return result
}

/*
func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}*/
