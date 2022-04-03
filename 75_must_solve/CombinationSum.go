package main

//39. Combination Sum
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var dfs func(i, total int, cur []int)
	n := len(candidates)
	dfs = func(i, total int, cur []int) {
		if target == total {
			curCopy := make([]int, len(cur))
			copy(curCopy, cur)
			res = append(res, curCopy)
			return
		}
		if total > target || i >= n {
			return
		}
		cur = append(cur, candidates[i])
		dfs(i, total+candidates[i], cur)
		cur = cur[0 : len(cur)-1]
		dfs(i+1, total, cur)
	}
	dfs(0, 0, []int{})
	return res
}

/*
func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{3, 5, 8}, 11))
}*/
