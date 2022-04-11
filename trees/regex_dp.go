package main

/*
class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        cache = {}

        def dfs(i, j):
            if (i, j) in cache:
                return cache[(i, j)]

            if i >= len(s) and j >= len(p):
                return True
            if j >= len(p):
                return False

            match = i < len(s) and (s[i] == p[j] or p[j] == '.')
            if (j + 1) < len(p) and p[j + 1] == '*':
                cache[(i, j)] = (dfs(i, j+2) or (match and dfs(i + 1, j)))
                return cache[(i, j)]
            if match:
                cache[(i, j)] = dfs(i + 1, j + 1)
                return cache[(i, j)]
            cache[(i, j)] = False
            return cache[(i, j)]

        return dfs(0, 0)
}
class Solution {
public:
    bool isMatch(string s, string p) {
        int m=s.size();
        int n=p.size();
        bool dp[m+1][n+1];
        memset(dp,false,sizeof(dp));
        dp[0][0]=true;

		//if star look two columns behind
        for(int i=2;i<=n;i++){
            if(p[i-1]=='*')
                dp[0][i]=dp[0][i-2];

        }
        for(int i=1;i<=m;i++){
            for(int j=1;j<=n;j++){
                if(s[i-1]==p[j-1] || p[j-1]=='.')
                    dp[i][j]=dp[i-1][j-1];
                else if( p[j-1]=='*'){
                    if(dp[i][j-2] == true)
                        dp[i][j] = dp[i][j-2];
                    else if(s[i-1]==p[j-2] || p[j-2]=='.')
                        dp[i][j]=dp[i-1][j] || dp[i][j];
                }

            }
        }
        return dp[m][n];
    }
};

*/

func isMatch(s string, p string) bool {

	m := len(s)
	n := len(p)
	if m == 0 {
		return (n == 0)
	}
	table := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		table[i] = make([]bool, n+1)
	}
	table[0][0] = true
	for j := 2; j < n+1; j++ {
		if p[j-1] == byte('*') {
			table[0][j] = table[0][j-2]
		}
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[j-1] == byte('.') || s[i-1] == p[j-1] {
				table[i][j] = table[i-1][j-1]
			} else if p[j-1] == byte('*') {
				if table[i][j-2] {
					table[i][j] = table[i][j-2]
				} else if s[i-1] == p[j-2] || p[j-2] == byte('.') {
					table[i][j] = table[i-1][j] || table[i][j]
				}
			}
		}
	}
	return table[m][n]
}
