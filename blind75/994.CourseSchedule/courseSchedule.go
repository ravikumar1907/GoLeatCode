package main

import (
	"fmt"
)

/*
207. Course Schedule
*/
// DFS based
func canFinish(numCourses int, prerequisites [][]int) bool {
	adjList := make([][]int, numCourses)
	for _, pre := range prerequisites {
		adjList[pre[1]] = append(adjList[pre[1]], pre[0])
	}
	visited := make([]int, numCourses)
	var dfs func(int) bool
	dfs = func(course int) bool {
		if visited[course] == 1 {
			return false
		}
		if visited[course] == 2 {
			return true
		}
		visited[course] = 1
		for _, nextCourse := range adjList[course] {
			if dfs(nextCourse) == false {
				return false
			}
		}
		visited[course] = 2
		return true
	}
	for i := 0; i < numCourses; i++ {
		if dfs(i) == false {
			return false
		}
	}
	return true
}

/*
Topological sort based BFS
*/
func canFinish2(numCourses int, prerequisites [][]int) bool {
	// Create adjacency list and in-degree array
	adjList := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	// Populate the adjacency list and in-degree
	for _, pre := range prerequisites {
		course := pre[0]
		prereq := pre[1]
		adjList[prereq] = append(adjList[prereq], course)
		inDegree[course]++
	}

	// Queue for BFS to process courses with in-degree 0
	queue := []int{}

	// Enqueue all courses with no prerequisites (in-degree 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Process the courses using BFS
	processed := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		processed++

		// Decrease in-degree for all dependent courses
		for _, nextCourse := range adjList[course] {
			inDegree[nextCourse]--
			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	// If all courses have been processed, there's no cycle
	return processed == numCourses
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(3, [][]int{{1, 0}, {2, 1}}))
	fmt.Println(canFinish(3, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish2(3, [][]int{{1, 0}, {2, 0}}))
}
