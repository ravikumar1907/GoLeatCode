package main

func spiralOrder(m [][]int) []int {
	var res []int
	l, r := 0, len(m[0])
	t, b := l, len(m)
	for l < r && t < b {
		for i := l; i < r; i++ {
			// top row fixed, move l-->r
			res = append(res, m[t][i])
		}
		// done top, increment
		t += 1
		for i := t; i < b; i++ {
			// right column (r-1) fixed, move t -> b
			res = append(res, m[i][r-1])
		}
		// done with right column, decrement
		r -= 1
		if !(l < r && t < b) {
			return res
		}
		// bottom row fixed (b-1), move r -> l
		for i := r - 1; i >= l; i-- {
			res = append(res, m[b-1][i])
		}
		// bottom (done move up
		b -= 1
		// left column l fixed, move bottom up b -->
		for i := b - 1; i >= t; i-- {
			res = append(res, m[i][l])
		}
		l += 1

	}
	return res
}

/*func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}
*/
