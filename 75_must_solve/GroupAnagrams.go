package main

import (
	"fmt"
	"strconv"
)

func calMapKey(str string) string {
	var bMap [26]int
	for _, c := range str {
		v := int(c - 'a')
		bMap[v] += 1
	}
	var s string
	s += "#"
	for _, v := range bMap {
		s += strconv.Itoa(v)
		s += "#"
	}
	fmt.Println(s)
	return s
}

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	hMap := make(map[string][]string)
	for i, _ := range strs {
		key := calMapKey(strs[i])
		hMap[key] = append(hMap[key], strs[i])
	}
	for _, v := range hMap {
		res = append(res, v)
	}
	return res
}

/*
func main() {
	//fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	//fmt.Println(groupAnagrams([]string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}))
	fmt.Println(groupAnagrams([]string{"bdddddddddd", "bbbbbbbbbbc"}))
}
*/
