package main

//3. Longest Substring Without Repeating Characters
// sliding window pattern
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	max := 0
	i := 0
	hMap := make(map[byte]int)
	for j := 0; j < n; j++ {
		if idx, ok := hMap[s[j]]; ok {
			if i < idx+1 {
				i = idx + 1
			}
		}
		hMap[s[j]] = j
		length := j - i + 1
		if max < length {
			max = length
		}
	}
	return max
}

//5. Longest Palindromic Substring
func longestPalindrome(s string) string {
	n := len(s)
	result := ""
	resultLen := 0
	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			if resultLen < (r - l + 1) {
				result = s[l : r+1]
				resultLen = r - l + 1
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			if resultLen < (r - l + 1) {
				result = s[l : r+1]
				resultLen = r - l + 1
			}
			l--
			r++
		}
	}
	return result
}

//Leetcode 647: Palindromic Substrings
func countPalindromSubstring(s string) int {
	/*res := 0
	n := len(s)
	var helper func(s string, l, r, n int) int
	for i := 0; i < n; i++ {
		res += helper(s, i, i, n)
		res += helper(s, i, i+1, n)
	}
	helper = func(s string, l, r, n int) int {
		c := 0
		for l >= 0 && r < n && s[l] == s[r] {
			c++
			l--
			r++
		}
		return c
	}*/
	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			res++
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	return res

}

//424. Longest Repeating Character Replacement
//Longest Substring with Same Letters after Replacement (hard)
func characterReplacement(s string, k int) int {
	res := 0
	l, r := 0, 0
	hMap := make(map[byte]int)
	n := len(s)
	var maxCount func() int
	maxCount = func() int {
		max := -1
		for _, v := range hMap {
			if max <= v {
				max = v
			}
		}
		return max
	}
	for r < n {
		hMap[s[r]]++
		for r-l+1-maxCount() > k { // shrink window
			hMap[s[l]]--
			l++

		}
		if res < r-l+1 {
			res = r - l + 1
		}
		r++
	}
	return res
}
func longest_substring_with_k_distinct(s string, k int) int {
	n := len(s)
	max := 0
	l, r := 0, 0
	hMap := make(map[byte]int)
	for ; r < n; r = r + 1 {
		hMap[s[r]]++
		for len(hMap) > k {
			hMap[s[l]]--
			if hMap[s[l]] == 0 {
				delete(hMap, s[l])
			}
			l++
		}
		if max < r-l+1 {
			max = r - l + 1
		}
	}
	return max
}

/*
Problem Statement
Given an array of characters where each character represents a fruit tree, you are given two baskets and your goal is to put maximum number of fruits in each basket. The only restriction is that each basket can have only one type of fruit.

You can start with any tree, but once you have started you can’t skip a tree. You will pick one fruit from each tree until you cannot, i.e., you will stop when you have to pick from a third fruit type.

Write a function to return the maximum number of fruits in both the baskets.

Example 1:

Input: Fruit=['A', 'B', 'C', 'A', 'C']
Output: 3
Explanation: We can put 2 'C' in one basket and one 'A' in the other from the subarray ['C', 'A', 'C']

Example 2:

Input: Fruit=['A', 'B', 'C', 'B', 'B', 'C']
Output: 5
Explanation: We can put 3 'B' in one basket and two 'C' in the other basket.
This can be done if we start with the second letter: ['B', 'C', 'B', 'B', 'C']
*/
func fruits_into_baskets(s string, k int) int {
	n := len(s)
	hMap := make(map[byte]int)
	l, r := 0, 0
	max := -1
	for ; r < n; r++ {
		hMap[s[r]]++
		for len(hMap) > k {
			hMap[s[l]]--
			if hMap[s[l]] == 0 {
				delete(hMap, s[l])
				l++
			}

		}
		if max < r-l+1 {
			max = r - l + 1
		}
	}
	return max
}

/*
Problem Statement
Given an array containing 0s and 1s, if you are allowed to replace no more than ‘k’ 0s with 1s, find the length of the longest contiguous subarray having all 1s.

Example 1:

Input: Array=[0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1], k=2
Output: 6
Explanation: Replace the '0' at index 5 and 8 to have the longest contiguous subarray of 1s having length 6.

Example 2:

Input: Array=[0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1], k=3
Output: 9
Explanation: Replace the '0' at index 6, 9, and 10 to have the longest contiguous subarray of 1s having length 9.
*/
func length_of_longest_ones(s string, k int) int {
	n := len(s)
	max := -1
	max_ones_count := 0
	l, r := 0, 0
	for ; r < n; r++ {
		if s[r] == '1' {
			max_ones_count++
		}
		if r-l+1-max_ones_count > k {
			if s[l] == '1' {
				max_ones_count--
			}
			l++
		}
		if max < r-l+1 {
			max = r - l + 1
		}
	}
	return max
}

/*func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(countPalindromSubstring("aaab"))
	fmt.Println(longest_substring_with_k_distinct("cbbebi", 3))
	fmt.Println(fruits_into_baskets("ABCAC", 2))
	fmt.Println(length_of_longest_ones("01100011011", 2))
	fmt.Println((maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})))
}*/
