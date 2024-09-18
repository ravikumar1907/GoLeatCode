package main

import "fmt"

/*
91. Decode Ways
https://leetcode.com/problems/decode-ways/description/
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
