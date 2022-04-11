package main

import "fmt"

// Course Schedule : LeetCode 207.
func canFinish(numCrs int, preReqs [][]int) bool {
	preMap := make(map[int][]int)
	for _, crs := range preReqs {
		preMap[crs[0]] = append(preMap[crs[0]], crs[1])
	}
	visited := make(map[int]struct{})
	var dfs func(int) bool
	dfs = func(crs int) bool {
		if _, ok := visited[crs]; ok {
			return false
		}
		prereq, ok := preMap[crs]
		if ok && len(prereq) == 0 {
			return true
		}
		visited[crs] = struct{}{}
		for _, pre := range prereq {
			if !dfs(pre) {
				return false
			}
		}
		delete(visited, crs)
		preMap[crs] = []int{}
		return true
	}
	for i := 0; i < numCrs; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}

// course
func courseOrder(numCrs int, preReqs [][]int) []int {
	preMap := make(map[int][]int)
	for _, pre := range preReqs {
		preMap[pre[0]] = append(preMap[pre[0]], pre[1])
	}
	visited := make(map[int]struct{})
	cycle := make(map[int]struct{})
	var output []int
	var dfs func(int) bool
	dfs = func(crs int) bool {
		if _, ok := cycle[crs]; ok {
			return false
		}
		if _, ok := visited[crs]; ok {
			return true
		}
		prereqs := preMap[crs]
		cycle[crs] = struct{}{}
		for _, pre := range prereqs {
			if dfs(pre) == false {
				return false
			}

		}
		delete(cycle, crs)
		visited[crs] = struct{}{}
		output = append(output, crs)
		return true
	}
	for i := 0; i < numCrs; i++ {
		if dfs(i) == false {
			return []int{}
		}
	}
	return output
}

func main() {
	preReqs := [][]int{
		{0, 1}, {0, 2}, {1, 3}, {1, 4}, {3, 4},
	}
	//fmt.Println(canFinish(5, preReqs))
	//fmt.Println(courseOrder(5, preReqs))b
	preReqs = [][]int{{0, 1}} //{0, 1}, {1, 2}, {2, 0},

	//fmt.Println(canFinish(3, preReqs))
	fmt.Println(courseOrder(2, preReqs))
	fmt.Println(courseOrder(2, [][]int{{1, 0}}))
	fmt.Println(courseOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}

//CanFinish(numscrs, prereqs):
// Initialize PreMap[0] = []int{}, PreMap[1] = []int{} ..... PreMap[n-1] = []int{}
// Set PreReq data for each course using input Prereqs
// Initialize Visisted Map
// 		dfs(crs):
// 		if crs in visisted return false
// 		if crs prereqs is empty return true
//  	visisted[crs] == struct{}{}
// 		for pre in prereqs of crs:
// 			if dfs(pre) == false return false
// 		finally return true
// for crs in numcrs:
//	if dfs(crs) == false return false
// return true
