package main

import "fmt"

/*
91. Decode Ways
https://leetcode.com/problems/decode-ways/description/
*/
/*
Dynamic Programming Relation: The fundamental relationship in this problem is:

dp[i]=dp[i - 1]+dp[i - 2]

so when i incremented by 1,
 dp[i-2] becomes dp[i-1] & dp[i-1] becomes dp[i]
 if we consider prev1 = dp[i-1] and prev2 = dp[i-2],
 when I increments prev2 = dp[i-1], which is prev1 so prev2 = prev1 and dp[i-1] becomes dp[i] which current and so prev1= current right

  // Base cases
    dp[0] = 1  // Empty string has 1 way to decode (doing nothing)
    dp[1] = 1  // First character is valid if it's not '0'

	base cases
	prev2 = 1 for empty string
	prev1 = 1 for first charracter
	 for starting i= 2
	 0 - 1 (s[i-2:i])
	 1 - 1 (s[i-1:i])

*/

func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}
	prev2, prev1 := 1, 1
	for i := 2; i <= n; i++ {
		current := 0
		oneDigit := s[i-1 : i]
		twoDigit := s[i-2 : i]
		if oneDigit >= "1" && oneDigit <= "9" {
			current += prev1
		}
		if twoDigit >= "10" && twoDigit <= "26" {
			current += prev2
		}
		prev2 = prev1
		prev1 = current
	}
	return prev1
}

func main() {
	ret := numDecodings("12")
	fmt.Println(ret)
	ret = numDecodings("226")
	fmt.Println(ret)
	ret = numDecodings("06")
	fmt.Println(ret)
}
