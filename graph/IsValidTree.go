package main

//261. Graph Valid Tree
/*
You have a graph of n nodes labeled from 0 to n - 1. You are given an integer n and a list of edges where edges[i] = [ai, bi] indicates that there is an undirected edge between nodes ai and bi in the graph.

Return true if the edges of the given graph make up a valid tree, and false otherwise.
*/

func validTree(n int, edges [][]int) bool {
	hMap := make(map[int][]int)
	visited := make(map[int]bool)
	for _, v := range edges {
		hMap[v[0]] = append(hMap[v[0]], v[1])
		hMap[v[1]] = append(hMap[v[1]], v[0])
	}
	var dfs func(i, prev int) bool
	dfs = func(i, prev int) bool {
		if visited[i] {
			return false
		}
		visited[i] = true
		edges, _ := hMap[i]
		for _, e := range edges {
			if e == prev {
				continue
			}
			if !dfs(e, i) {
				return false
			}

		}
		return true
	}
	return dfs(0, -1) && (len(visited) == n)
}
