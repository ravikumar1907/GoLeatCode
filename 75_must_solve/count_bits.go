package main

//Leetcode - 338 Counting Bits
/*
0	0000  - d[0] = 0
1	0001  = dp[1] = 1
2	0010    dp[2] = 1
3	0011    dp[3] = 1
from 4 most sig bit moved 3rd position and after 3rd position pattern of 0 -3 repeats from 4 - 7
4       0100    dp[4] = 1 + dp[4 - 4]
5       0101    dp[3] = 1 + dp[5-4]
6       0110            1 + dp[6-4]
7       0101            1 + dp[7-4]
8       1000            1 + dp[8 - 8]
 offset = 1
  if offset * 2 = i
     offset = i

*/
func countBits(n int) []int {
	ans := make([]int, n+1)
	offset := 1
	ans[0] = 0
	for i := 1; i < n+1; i++ {
		if offset*2 == i {
			offset = i
		}
		ans[i] = 1 + ans[i-offset]
	}
	return ans
}
