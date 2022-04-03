package main

/*
Number of Connected Components in an Undirected Graph
Medium
Robot Return to Origin
Easy
Sentence Similarity
Easy
Sentence Similarity II
Medium
The Earliest Moment When Everyone Become Friends
Medium
Detonate the Maximum Bombs
*/
/*
class Solution {
	public:
		vector<int> parent;
		vector<int> rank;

		int find(int n) {
			return (n == parent[n]) ? n : parent[n] = find(parent[n]);
		}

		void unionn(int a, int b) {
			a = find(a), b = find(b);
			if (rank[a] > rank[b])
				swap(a, b);
			rank[b] += rank[a];
			parent[a] = b;
		}

		int findCircleNum(vector<vector<int>>& isConnected) {
			int n = isConnected.size();
			int provinces = n;
			parent.resize(n);
			rank.resize(n, 1);
			for (int i = 0; i < n; i++)
				parent[i] = i;

			for (int i = 0; i < n; i++)
				for (int j = 0; j < n; j++)
					if (isConnected[i][j] == 1)
						if (find(i) != find(j))
							unionn(i, j), provinces--;

			return provinces;
		}
	};
*/
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	provinces := n
	var find func(n int) int
	find = func(n int) int {
		res := n
		for res != parent[res] {
			parent[res] = parent[parent[res]]
			res = parent[res]
		}
		return res
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && isConnected[i][j] == 1 {
				p1 := find(i)
				p2 := find(j)
				if p1 != p2 {
					provinces--
					if rank[p1] > rank[p2] {
						parent[p2] = p1
						rank[p1]++
					} else {
						parent[p1] = p2
						rank[p2]++
					}
				}
			}

		}
	}
	return provinces
}

func findCircleNum1(isConnected [][]int) int {
	n := len(isConnected)
	provinces := 0
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= n {
			return
		}
		if visited[i][j] || isConnected[i][j] == 0 {
			return
		}
		visited[i][j] = true
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && isConnected[i][j] == 1 {
				dfs(i, j)
				provinces++
			}
		}
	}
	return provinces
}

/*func main() {
	m := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(findCircleNum(m))
	fmt.Println(findCircleNum1(m))
}*/
