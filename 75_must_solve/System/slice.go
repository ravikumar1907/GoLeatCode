package main

import "fmt"

func main() {
	var s1 []int
	s2 := make([]int, 0)
	s3 := []int{1}
	s1 = append(s1, s3...)
	s2 = append(s2, s3...)
	fmt.Println(s1, s2)
}
