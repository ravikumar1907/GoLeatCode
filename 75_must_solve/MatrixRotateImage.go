package main

func rotate(m [][]int) {
	l, r := 0, len(m)-1
	for l < r {

		for i := 0; i < r-l; i++ {
			t, b := l, r
			tmp := m[t][l+i]
			m[t][l+i] = m[b-i][l]
			m[b-i][l] = m[b][r-i]
			m[b][r-i] = m[t+i][r]
			m[t+i][r] = tmp
		}
		l++
		r--
		/*
			          l=0, r = 2
			          t=0, b = 2
			          tmp := m[t][l+i] ==> m[0][0]
			            m[t][l+i] = m[b-i][l]  ===> m[0][0] = m[2][0]
			            m[b-i][l] = m[b][r-i]  ===> m[2][0] = m[2][2]
			            m[b][r-i] = m[t+i][r]  ===> m[2][2] = m[0][2]
						m[t+i][r] = tmp   =========> m[0][2] = tmp

			                tmp = m[0][0]
			                m[0][0] = m[2][0]
			                m[2][0] = m[2][2]
			                m[2][2] = m[0][2]
			                m[0][2] = tmp
		*/

	}
}
