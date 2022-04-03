package main

func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	n := len(intervals)
	for i := 0; i < n; i++ {
		if newInterval[1] < intervals[i][0] {
			result = append(result, newInterval)
			return append(result, intervals[i:]...)
		} else if newInterval[0] > intervals[i][1] {
			result = append(result, intervals[i])
		} else {
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
		}
	}
	return append(result, newInterval)
}

/*
func main() {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newI := []int{4, 8}
	fmt.Println(insert(intervals, newI))
}*/
