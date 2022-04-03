package main

import (
	"fmt"
	"strconv"
)

// 659: Encode Decode strings
func Encode(strs []string) string {
	var res string
	for _, s := range strs {
		res += fmt.Sprintf("%d#", len(s)) + s
	}
	return res
}

func Decode(str string) []string {
	var res []string
	i := 0
	r := []rune(str)
	n := len(r)

	for i < n {
		j := i
		for r[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(string(r[i:j]))
		j++
		s := r[j : j+length]
		res = append(res, string(s))
		i = j + length
	}
	return res
}

/*func main() {
	strs := []string{"Neet", "Code"}
	ret := Encode(strs)
	fmt.Println(ret)
	fmt.Println(Decode(ret))
}*/
