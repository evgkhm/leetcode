package main

import (
	"fmt"
)

func main() {
	//answerKey := "TTFF"
	//k := 2

	//answerKey := "TFFT"
	//k := 1

	//answerKey := "TTFTTFTT"
	//k := 1

	//answerKey := "TTFF"
	//k := 2

	answerKey := "FFFTTFTTFT"
	k := 3

	//answerKey := "FFTFTTTFFF"
	//k := 1

	//answerKey := "FFFTTFTTFT"
	//k := 3

	//answerKey := "FTFFTFTFTTFTTFTTFFTTFFTTTTTFTTTFTFFTTFFFFFTTTTFTTTTTTTTTFTTFFTTFTFFTTTFFFFFTTTFFTTTTFTFTFFTTFTTTTTTF"
	//k := 32

	//answerKey := "TTFFTFTFFFTTTFTFFTTFTFTTFTTTTTFFFTFTTFTTFFTTTTTFTFFTTFTFTTFFFFTTTFTTFTTFTFTTTFTTTTTFTFTFFTTTTFFFFTFTTTFTFFFTFTTTFFTTFTTFTTFTTFTTTFFFFTFTTTFFFTTTFFTFFTTFFFFTFTTTTTFTTFFFFFFFFTTFFTTFFFFTTFTTTFFTTFFTTFTFTTFTFFFFFTTTTFFFFTTTTTTTFTTFFTTTTTTTFTFFFTTTTTTFTFTTTFTTTFTFFTFTFFTFFFFTTTTFTTFTTTFTTTFTFFFTFTTTFTTFTFTFFFTFFFFFFFFFTTTTTFTTTTTTFFTTFTTFTFFFTFTFTFFTTFFTTFTTTTFTTFTTFFTFFTTTFTTTFFFTTTTTTFTTTTTTFTFTTTFFTFFFFTFFFTFTTFFTFFTTTFTTFFTFFTFTFFTFFTTFTTFTFFTFFFFFTTFTFFFFFFFTTTTFTTFTFFFTTTFFTFTFFFTTTTFFFFTFFTTFFFTTTFTTTFTTFTFTFFTTFFTFFFTFTTTTTTFFTFTTTTTFTFFTTFTTFTTTTTTTFTTTFTTTFFTTTFTFFTTFTFTFTFFFFTTTTFFTFTTTFTFTFTFTTFTFTFTTFTTTFFFFTTFFTTFFTTTFTFFFTTTFFTTTFTTTFTFFFFFFFTFFFTFTTFTTFTTFFTFTTFTTTFTFTFTTTTFFTTTTFTFTTFTTTTFTFFFFFFTTFTFTTTFTTFFFFTFFTFFTFFTTTFFFTTTFFTFFFTFTFTTFTTFFFFFFFFTTFTFTTFTTFTFFFFTFTTTFFFTTTFFTFFFTFTTTFFFFTFTFTFFFFTTFFTTFTTFTFFFTFTTFFTFTFFTTTFFTFTFTTFTFTFFFFTFTFFFTFFFFFFTFTTTFTTFTFTFTTTTTTFTTTFFFFTFTFTTFTFTFFFTTFTTFTTTTTTFTTFTFFFTTFTTFFTTTFTFTFTTFTTTFFTFFFTTFTFFFTTTTFFFTTFFTTTTFTFTTTTFTFTFFTTFFFFFTFTFFTFFFTFTFFTFFFFTTTTTTTTTTFFTTFFTTTTFTTTTTTTFTTTFFFTTTFFFTTFFTTTTTFTFTFTTTFTTTFFTTTFFFFTTTTTFFTTFFFFFTTTFTFTTFFTTTFFTTTFFTFTFFFFTFFFTTTFTFTTTTFFTFTTTTTFTFFTFTFTTTTFTTFFFTFFTFFTTTTFFTTTFFFTFFFTTTTFFTFTTFFFFFTFFTTFFFTFTTFFFTTFFTFFTFTTTTFTTTFFTFFTFTFTTTTTTFTFTTTFFTFTTFFTFFTFFTTTFFTTTTTTTTFTTTTFFTTTFTTFTTFFTFFFFTFTFFTFFTFTTFTFTTFFTFFTTTTTFTTFFTFFTTTFFFFTTFTFFFFFFTFTFFTTFFTFFFTFTTFTTFTTTFTFFFFTFTTFFFFTFFFTTFFFTFTFTFFTTFFTTTFTFFFFTTFFTFTTTFTFFTFFFTFTFFFTFTFFTFFFFFFTFTFTFFTFTTTFFTTTFFTFFTFTFFFTFTTTFFFTTTFFFTTFFTTTTFFFFTTFFFTFTFFTFTFTFFTFFTFTTTFTFFTFFTFFTTTTFTTFTTTTFTTFFTTFTFTFFFTFFTTTTTTFFTTFFFFTFTTFFTTFFTFTFTTTTFTFFTFFFTTFTTFFTTFTTFTTFTTFFTFTTTTTFTFTTFTFFFTTFTFTFFTTFTFTFTTFFFTTTTTTTFTTFTFFFTFFTFFFTTTTFTFTFTFTTTTFFTTTFFFTTFFTTFTFFTFTFFTFFTFFFTTFFTTFFFTFFFFTTTTFTFFFFFTTFFFTTTFTFTTFTTFFFTFFTTFFTTTFFFFFFTFFTFTTFTTFTTFTTTTTFTTTTFFTFFTTFTTFTTFTTTFFTFTTFFFFTTFTFTTTTFTTFFFTFFTTFFTTTFFTTFFFTFFFTFTFTFFFTTFTTTFFFTTTFFFTTTFFTTFTTTTFTTFTTFTTFTFFFFFTFFFTFFTTFTTFFFTTFTFTFTTFTFFFTFTTFTFFTTFFFFFFTTFTFFTTFTTTFFFTFTTTTTTTFFTTTFFFTFTTFFTFTFFFTTTFFTTTTFFFTTFFTFTTTTTFTTTTTFFTTFFFTFTFTTTTFTFFTTTFTFTFFFFFFFTTTTFTFFTFTFFFFFTFTTFFTFTTTFTFTTFFFTTFFTFTFTTFFTFTFTTTFFTTTTTTTTTTTFTTFFTTFFFFTTTFTTFTTFFFFFFFTTFFFFFTTTTTFFFTFFFFFTTFTTFFTFFFTTFTFTFFTFTFFTFFTTTFTFFTTFTTTFTTFFTFFTFFTFTTFFTTTFFFFTFFFTTFFTFFTTFFFFTFFFTTTFFFFFTFFTTTFFTFTTTTTTFFFTTTTFTTTFFFTTTFTTFFFTFTFFTTTFTTTTTTTFTTFTTTFFTTFFTTTFFFTFFFTFTFFFTTTTTTFFTTTTFTTFFFFTTTFTFFFFTFTTFTTFTFFTFTFTTFFTTFFTFFFTFFFTFTFTTTTFTFFFFTFFFFFTTTFFTFFTFFTFTTFFFTTTTFFTTTTTTFFFTFTTTTFFTTTFFFTTFFFFTTTTFFTFFTFTFTTFFFFFFFTFFFFFTTTFTTFTTTTFTFFTTTFFFFFFFFFFFFFFFTTFFFFTTTFFFFTFTFTFFFFFTTTFFFTFFTTTTFFTTFFTTFTTFTTFFTTFTTTFTTTFTTFTFTFFFTFFTFTTFFFTFFTFTFFFFTFFTFTFTTFTTFFTTTTFTFFTFTFTTTTFFFFFFFFTFTFFTTTTFFFTFTTTTFTFFFTFFFFTTFTTFFTTTFTTFTFTFFFFFTFTTTFTTFTFTTTTFTTFFFFTFFFFFTFFFFFTTTFTFTFTFFTFFFTFTFFFTTTTFFTTTFTTTTFTTTFTTTFTFFTTTFTTFTTFTTFFFFTTFTFFFFFTTFFFTFFTTTTTFFFTFFTFTTFTFFFTFFTTTFTFFTTTTFTTFFTFFTTFFFTTTFFFFFFFFFTTFFFTTFTTFFTFTTFFFFTTFTTFTTTTFFFTTTTFTFFTFFTFFTFTFFFFFFTTFFFTFFTFTFTFTTTTTFTFTTFFFFTFTFFFTTTFFTTTTFFTFTTTFFFFTFFFFTTTFFTFTTTTTTFTTTFFTFFFTTFFTFTFFFTFFTTTTFTFFFFFFTFTTFTFTFTTFFFFTTFFFFTFTTTFTTFFTTFFFFFTTTTFFTTTFFFFTFFFTTFTTFFFFTTTTFTFFTFFFFTTTFTTTFTTFTTFTTFTFFFTTTFTFTTFFFFTFFTFTTFTTTFFTTTFTTTTTTFFFTFTTFTTFTTFFFTTFTFFTTTFFTTTTTTFTTFTTTTTFFTTTTTTTFFTTTTTTFTFTTTTFFFFFTTFTFFFFFTTFFFTFTTFTFFFTFFTFTFFTFTFFFFTFFFTFTTFFFTFFFTTFFTFFFFFFTFTFTFTTFTFFTFTFFTTFFTTFFTFFFFTTFTFTFFTTFTFTTTFTFTTTTFTFTFTFTTTFFFTTFFTTTTFFTFFFTFFTTFFFFTTTFTFTTFFTTTFFFFTFFTFTTFTTTTFTTFTTFTFTTTTTFFTFFFTFTFTFTTTFFTTTFTTFTFTTTFTTTFFTFTTTTFTTTFFFFTTTFTTFTFTFTTFTTFFTTFTTTTFTTFFTTTFTTTFFFTFTTFFFFFFFTTTTTFFTFTFTTTTFFTTTTFTFTFFTFTFFFFTFTFFTTTFFFTFFFFFFTTTFTFFFFFTFFFTTFTFTFFTFFFTFFTFTTTFFTTTFFTFTFTTTFFTFTFTTTFFTTTTFTTTTTTTTFTFFFTTFFTTFTTFFTTFTTTFTFTFFTTTTFFTFFTFTTTTTTFTFFTFTFTTFFFFFFTTTTTTFTFFTTFTTTTFFTFTFFFFFFFTFTTFFTFFTFTFTTFFFFTTFFFFTTFFTFFTTFTFTTFTTFTFFTFFFFFFTFTFTTTFTTFTTFTFTFTFFTFTFFFFFTFTFFTTFTFFTTTTTFFFTTTFTFTTFTTFFTTFTTFFTTTTFTFFTTFTTFFTFTFTTFTTTTTTFFTTTTFTTTTTTFTFFTTTTTTFFTFTTFFFFTFFFTFTTTFFTTTTFFFTFFFFFFTFTTFFTTTFFFFFFTFTFFTTTTTFFTFTTTTTTTFFFTFFFTFTTTTFTFTTTFFFFTFTTFTFTTTTTFFTTTTTTFFFFFTTFFTFTFFFTFTTTTTFTFFTFFFTFFTFTTFTFFFTTFFFFFFFFFFFTTFFFTTFFTFTFTFFFFTTTFTFFFTTFTFTFFTTTTFFFTTTTFFFTFFFFTFTFTFTTTTTFFTTTFTTFFTFFFFTFTFTTFTTTTTTTTFFTFFTTTFFFTTFFTTFTFFFFTTFFFTTTTFFTTFTTTTFFFTFFTTTTFFTTFTFTTFFFFTTFTTFTFFFFFTTTTTFFFFFTFFFTTTTFTFTFTFTTFFFFTTFFFFTTTFTFTFTFTFTFFTFFFFTFTFTTTFTFFTTFFFTTFFFFFTTTTTFTTFTFFFTFFTFTTFTFFFTTTFTTFTTFTTFFFFTTFTFTFFTTTTFTTTTTFTTFFFTTFFFFTFTFTFFFFFFFFFTFTTFFTTFTTFTFFTFTTTTTFTTFTFFFTFTTFFTTFTTTTFTFTTTTFFTFFFFTTFFTFFTTTFFFTFFTFFFTTTTTFTFFFTTFFFFTFTFTTTTTFTTFFTFTTFFTTFTFTTTTFFTFTTTTTTTTTFTTFTTTFFTFTTTTFTFFFTTFFTTTTTFTFFFFFTFTTTTFTTTFFTFFTTTFTFTFFFTTFFFFTFFFTTFTFFTTFTFFTTTFFFFTTFFTFFTTTFFTFTTFFFFFTFFFFFTTFTFTTFTFTFFFTTTFFFTFTTFFTTFTTTFFTFFFFTTTTTTTTTFTTFTFFTTFFTTTTTFTTTFTTTFTTTFTFFTFFFTFFTFFTFFFFFFFTFFTFFFFTFFTFFFFFTFTFTFFTFFTFFFFTTTFFFFTFTTFTFFTFFTFTFFTFTTFFFTTTFTFTTTFFTFTFFFTTTTFTFFTFTTTFFTTTFTFTTTFFTTFFTTTTTTFTFTFFTFFTFFTFTTFFTTFFFFTFFTTTFTTFFFFFTFTFTFTTFFTTFFFTTTFFFFTFFTTFTFTFFTTTFTTFTTTFTFFTFFTFTTFFTTTTFTTFFTFTTTFTFFFTTTTTTFTFFFFFFTTTFFTFFTTTFTFFTTTTFFTTFTFTTFTFFTFTTTFTTTFTTTTTFTFFTFFFTTFTFFTFFFTFTTFFFTTTTTFTFTTFFTTFFFFTTTFTTFFFFFTTTFTTTTTFTTTFTTTFFTTFTTTTTTFTTTTTFTFTFTTTFTFTFFFFFTTTTFTTFTTFTFTFTFFTFTTTFTFTTTFFFFFFFTFFFFTTFFFTTFFTFTFTFFTTFFFTTFFFTFTTFFTTTTFFTFFFFTFFTTTFTTFFTFTTFTFTTFFTFFFTFFFFFFTTFTTTFTFTFTTFTTTTFTFFTFTTFTFTTTFTTTFFTTTTTFFTFFFFFFFTFTTTFTTFFFFFTTTFTFTFFTTFTTTFFFFTFTFFFTFTTTTFTTTTTFTFFFTTTFTTFTTTFFFTTFTTFTFFFFTFFTTTFFTTTFFTFFTTFTFFFTTTTFTTFTFFFTTFFTTFFTTTTFFFTFFFFTFTTTFFTFTFFTTTTFTFFTFTTFTFTTFTTTFTTTFTTTFFFTTTFFFTFFTTTFTFTFFFFTTFTFFFTFTFFFTTTFTFFTFFTTTFTTTTFFFFTTTTTFFFFFTTFTTFTFTTTTTFTFTFTFFTTFFTTFTTFTFFTFFFTFFFTFTTFTFFTFFTFTFTFFFTTFFFFTFTFFFTFFTTTTFFTFTFFFFFFFFFTFTFTFTTFFTTTTFFTFFFFFFFTFTFTTFFFFTTFTFFFTFFTFFFTFFFTTFTFTFTFFTFTTFTTFFFTFFTFTFTFTFTFFTTTTTTFFFTTTFFTFFFFFTTFFFTTTFTFFFTFFFTTTTTTTFFFTFTFTTFFFFFFFTFTFTFFFTFFFFFFFTFFTFTFFFTTFFFFFFTFTTFFFFTFFFTFTTFFTTTTTTFTFFFFTTFFFTFFFTFTTTFFTTTFFTTTTTFFFTTFFTTTFTFFFFTFFFFFTFFTFFFFFTTFTFFFFTTFTTTTTTFFFTFFFTFTTTFTFTTTFFTFFFFTTTFFFTFFFFTTTFFFTTTTFTFTFFFFFFTTTTFTTTTFFTFFFTFTFTTFFFTFTTFFFTFFFTTTTTFTTTTTFFTFTTFTTFFTTTTTFTFTFTTFFTTTTFFFTTTFTTTTFFTFFTTTFFFFFFFFTFFFTTTTFFFTTFFTTTFTFFTTTTTFTTFTFTTTFFFTFFTTTFFTTTFTTFFTTFTTTTTTFFFFTFTFTTFTFFTTFTTFTFTFTTTFTFFFFFFFFTTFTFFFFTTTFFFFTTFFFFFTTTFTFTTTTTTFTFFFFTTTTTFTFFTTFFTTFFTFTFFTTFTFFTTFTTFFTFFFFFTTFTTTTTTTTFTTFFFFFFFFTFFFTFFFFFFTTTFTTFFFFTFFFFTTFFTFFFTTFFFFTFFFFTFFFTFFFTTTTFTTTFTTTFFTTFTTFTFTFFFTFFTTTFTTTFTTFTFTFFFFTFFFFFTFTTFTFFFFTFTTFTFFTTTFFFTFTTTFFTTTFFFFTTTFTFFTFFFFFFTFTFFTTTFFFFTTFFTFTFTTFFTTFFFFFFFFFTFTFFFFTFTFFFTFFTTTFTFFTFTFTFFFTTTTTFFFFTFFTTFFTTFFFFFFTTFFFFTFFFFTFTTTFFFTFFFFFTFFTTFTTFTTTFFTFTFTFFFTFFFTFTFFFTTFFFFTTFTTFTTTFFTTTFTFFFTFFTFFTFTTFTFFFTFTFFTFFFTFTFTTTFTTTFTTTTFTTFFTFFTFFTFFFFTFTTTTFTFFFTFTTFTFTFFTTFFTTTTTTTFTFTFTTFTFFFFFFFFFTTFFFFTTTFTTFFTTFTTTTTFFFTFFTFFTTTTFTTTFFTTTTFTTFTTTFFFFFTTFFFTTFFTFTFFFTFTTFTTFFTFTFTTTTFFFFFFFFTTFTFTTFFTFTTFTFTFTTTFTTFTFFTFFTTFTTTFTFTFTTTTFTFFFTTTTFTFTTTFFFTFTFFFFFTFTTTTFTTTTTTFFFTTFFTFFTFTTTFFFTFTTFFTFTFFTFTTFFFTFFFFFTTFFTFFFTTFTTTTTFFFFFFFTTFFTFFTTTFFTFFTFTFFFFFTTTTFFFFFFFFFFFTFTTFTFFFTTTFFFFTFFTFTTTFFTFFFFFFTTFTFTFFTFTTTFTFTTFTTTFTFTTTTTTTTTFFFTFTFFTTFFTTFFFFTFTTFTFTTFFFFFFFTTTTFTFFFTTFTFFFTFFTFFFFTTTTTTTTTFTTFFFTTFFTFTTFTFFFTTFFFFTFTTTTTFTTFTTTFFFTTTTTTTFFFFFTFFFTTTFFFTFFTTFTFFFFTTTFTFFFTTTFFTTTFTFTTFFFTTTFFTFFTFTTFTTFFFTTTFTTTTTFTTTTFTFTFFTTFFFTTFTFTFTTFFFFTFTFFTFTFTFTTTTFFFTFFFTFFTFTFFFFTFFFTFFFTTTTTFTFTFFFTTTFTTFFFTFFTFTFFTTFFTTTFTTTFTTFTFTTTFTTFFTFFFFTTTFTTTFFFFTTTFFFFFTFTTTTFTFFTFFTTTTTTTFFFFTFFTFFTFFFTFTTTFFTTFFFFTTFTFTFFTTFTFTFFTTTTTTFTFFTTTTTFFTTTFFFTFTFTTTTFTFTFTTTTFTTFTFFFFTFTFTTFFFFTTTFFTTFTTFTTTFTFTFFTFTFFFFFTFFTTTFFTTFTFTTFFTFTTTFFTFTFTTFFFFFFTFTTFTTTTTTFFTTTTFTTTFFTFTTFFFTFFTFTFFFTFFTTTFTFTTTFFTTFTFFFTFTTTTFFFTFFFTFFTFFFFTFFFFFTFFTTFFFFTFFFTFTTTFTTTTTFFFFTTFFFTTTTFFFTTFTFFFTTTTFTTFFFFTTFFTTTFFFTFTFFTFTTTFTTTTTTFTFTFTTFFTFFFFFFTFTFTFTTFTFFFTTTTTTTTFTFFFFTTFTFTFTTTFTFFFTTTTTFFFFTFTFTFTFFFFFFTTTTFTFFTTTFTFTFFTTTTTTFFTFTTFTFTTFFTTTTFFTFFFTFTTTFTFTTTTTTFTFTTTFTFTFFFTFFTFTFFTFFTTTTTFFFFFFTTFFFFTFFTTTFFFFTTTTTFFTTFFTTFTTTFFFTTTFFFTTTTTFFFFTTFTFTFFTTFFFFFFFFFTFTTFTTTFFTTFFTFFFTTTFTTTFTFTTFTTTFFFTFFFFFTTTFFFFTTTFTTTFTTTFTFFFTFTFFTTFTFFTFFTTFFFTTFFFTFTTFTTFFFTFTFFTFTFFTTTFTFTFTFFFFFTFTTFFFFTTTFTFTFFTFFTTFTTFTTTTTTTFFTFTFTFFTTTTTTTTFFTFFTTTTTFFTFFTFFFFTTTTTTTFTTFFTFFTTTTTFFTFTFTFTFFFFTTTTTTTFTTFFTTFTFFFTTFTFFFFFTFFTTTFTTTFFTTFFFTFFFTFTFFTTFTTTFFTFTTFTFFTTTTFFTTTFFFFTFTFFFFTTFFFTTTFTFFTTFTFFFFTTTFFTTTTTTFFTTTTFTFTTTTFTTTFTFFFFTTTFTTTFTFTFFTFTFFTFFTTFFFTTTFTTTFFTTFTTTTTFTFTTTFTTTTTFTFTFFTTFTTTFFTTFFFFFTTTTFFFFTTFTFFFFTTFFFTTFTFFFTFTTTFTTFTFTFFTTTFTFFTTTTTFTFFFTFFTFFFTFTFTTTTFTFFFTFFTFTTTFFTTTFFTFTFTFTTTFFFFTTFFFFFFFTTTFTTFFFFTFFFTTFTFFTTFTTTTTFFTTTFFFTTTTFFFTTTFTFTFFTTTFFTFFTFTFFTTTTFTFFFFFFFFFTTFFTFTFFFTFFFTFTFTFTFTFFFFFTFFFTFFFTFTFTTFTFFFTTFTTFTFTTTTFFTTTFFTFTTTTTFFTTFTFFTFFTFTTTTTFFFTFFFTTFFFFFTTTTFTFFTTFFFFFFFTTTTTTFTFFTFFTTTTTTFFFFFFTTFFFFFFTTFTFFTFFFFTTTFFTTFFTTTTFFTFTTTTTFFFTFTFTFFTTFTTTTTTFFTFFFTTFTTTTTTTFTTTFTTFTTTFTTTTTFFFTFTFFFFFTTTFTTTTTTFTFFFFFFTFFFTFTFFFTFFTFFTFFFFFFTFTFFFFFFTFTTFTFTFTTTFTFFTFTTFTFTFTFFFTTTFFFFFTTFTTTFTFTFFTFTTTFTFTTFTFFTFTTFFTTFFTFTTFFFFFTFFTTFTTTTTFTTFFTTTFTTFFFTTTTTTTFFFFFFTFFFFFTFTFFTTFTFFTFTFFTFTTTFTFFFFFTFTFFFTFFTFFFTFFTTTFFTFTFFFFFTFTTFFTFTFTTTTTFTFFFFTTTTFFTTTFTFTTTTFFTFFTTFFFTFTTTFFFFFFFTFTTFTFFFTTFFTFTFFTFFTTFFFFFFFFFFFFFTTFTTTTFTFFTFFTTFFTTTTFFFFTFFFTTFFFTTTTFTFFFTFTTFTTTFFTTTTFTFTFTFFFFFTTTTTTTFTFTFFTFFFFFFTFFFFTTTTFTTTFFTFFFFTTFFTFTTTTTTFTTTTTTFFFFTFFFTFFTFTFFTTTTFFTFFTTFTTTFFFFTFFTTTFFTTFFTFFFTTTFTTTFFTFTFTTFFFFFTFFTFTTFFTFFFFFTFFFTFFFFFTTTTTTFFFFFTTFTTFFFTFFTFFTFTTTTTTTTTTFTFTTTFFTFTTTTFTTFFTFFFTTFFTFTFFFFFFFFFFFFFTFFTFTTFTFFFTTFTFTFTTFTTFFTTTTFFFFFFTTTTTFFFFTTTFTTFTTTTFFTTTFTFTFTTTTFTFFFTTFFFTTTTFFFTTTTTFTFFFTTTFTFFFTTFTFTTTTTTTTTTTTFFTTTTFFFFTTTTTTTTTFFTTFTTFFFTFTFTTTTTFFTFTTTTTTTFFFFTFTFFFFTFFFTFFTTTFTTFTTFTTTFFTTFTTFFFTTFFFFFTTTTFTTFTFTTFTFTFTTTTFTTFFFFFTTTFFTFTTFTFFFFFTTTFTTFTTFFTFTFTTTTTFTFFFFFFFFFFFTFTFTTTFTTFFFTFFFTTFFFTTTFFTFFTTFTFTTFFTTFFTTFTTTFTFFFFFFFTFTFFTFFTFTTFFFFFFFFTTTTTTFTTFTTFFFTTFTFTFTFFTTFTTFFTTTTFFTTFFFTFFTTFTFFFTFTFFFFFFFTTFFFTFFTTTTFTFTFFFFFTTFTFTTTFFFTTTTFTFTTFFFFTTTFFFTFTTFFTTTFTFFFFTTTTFFFFFFTTFTTFTFTTFTTFTTFFFFFTTTTFFFTFFTFTTFFFFFTTFTTFTTFTTTTTTFFTFFFTTFFTTFFTFFFTTTTTFTTFTTFTFFTTTFTTFFTFFFFFTFFTFFFFFFFTFFFFFTTTFFFFFTTTFFTTTTFTFTTTTFFFTFFTTTTFFTFTFTFFTTFFFFFTTFTFTTFTFFTTTFFTTFTTFFTTFTTFTTFTTTTTFTTFTTFTTFTFTTTTTTFFFTTFFTFFTFTFTFFFTFFFFFFTTFFTFTFFTFTTFTFFTFTTTTFFTTTFTTTTFFFFTTTFTFFFFTFFTFFFFTFFTTFTTTFTFTFFTTTFFFTFFTFTFFFTTFFFFFFFTTFTFTTTTTTTFTFFTTFFFTTTTFFFTTFFTTTTTFFFTFTTFFFFTFFFFTFFTTTFTFFFTTFFFTFFFFFFTTFTFTTFFFFTTFFFFTTFFTFFTFFTTFTFTTTTFFFTFTFTFFFFTTTTFFTTFTTTTFFFFFFTFTTFFFFFFFTTFTTTTFTTFFFFTTTFTFTTTFTFFTTTFFTTTFFTTTTTTFFFTTFTFTFTTFFTFFTTTFFTFFTFTTFTTTFTFTFTTTTTTTFFTTTFTTTFTFFFFFTFFTFFTTTFTTTTFFTFFFTTTTFTTTTTTTTTTFFFTTFTTTFFFTTFTFTFFTFTTTTTTFTTFFTTFFFTTFTFTTFFFFFFTFTFTFTFFFFFTFTFFTFTFTTFFFFFTFTFFFFTFTTTFFFTFTFFFFTTFTFFTTFFTFFTTFTFFFFTTFFFTFFFTTTFFTTTTFTFTTFTFFTTFFFTFTTFTFTTFTFTTTTFFFTFTFTTTTTTTTFFTTFTFTFTFTFFTFTTTTFTTTTFTTTTFTFTFFFFTFFFFFFTFTFTTFTFFFTTFTFTTFTTFTTFFFTTFTTFFFFFFFTTFFFFTTTFFTTTFFTTFTFTTTFFTFFTTTTFTFFFTFFFFFFTTTFTTFTTTFFTFFTTTTTFTFFFFFFTTFFTTTFTFFTFFTTFFTTTTTTTTFFTFTFFTTFTFFFFFFFTFTTFFFFFTTFFTTTTTTTTFTFFFTFFTTTTFFFTTFFFTFFTTTTFFTFFTTFFFTTFTTTFTFTFTFTFFTTFFFFFFTTTFTTTFTFFFTFFFFFFFFTFFTTFTFFTFTTFTFFTTTTFFTFTTTTTTTFFTFFTTFFFTTFTTTFTFFFTFTFFTTFFFFTFFTTFTTTTFTFFTTTFFFTFFFTFFFTTFFTTTFFTFFTTFTFFFFTTFFTFFFFTTFFFFTTFTTFFFFFFFFFFFFFFTFFTFFFFTTFTFTFFTTFFFTFFTTFFFTFTFTTTFFTTTTFFFTFTFTFTTFTFTFTTFTFTTTFFTFFTTFFFFFFFFTTTTFTFFTFTTFTTTTTFFTFTFTFFFFFFTTTTFTFFTFFFFFTFTFTFTTFTFFFFFTFTTTTTFTTTTTFFFFFTTTFFFTFTFTTFFTTFFTFTFTFFFTTTFFFTFFFTFFTFTTTTTFTFFTTFFFFFFTFTTTTTFFFTFTFTTFTTTTFFTTTTTFFFFTTTFTFFTTFTTTTTFTTFTTFTTTTTFTTFTTTTFFFFTFFFFTTFFFTTFTTFTTTFFTFTFTFFFFTFFTTTTFFFFFFTFTFFTFTTFFFFTTFTFFFFFTFFTTTTFFFFFFFFFFTFFTTTFTFTFFTTFTFFFFFFFFTFFTTFFFTTFFTFTTFFFFTFFTTTFTFFFTTFFTTFTFTTFFFFTFFFFFFTTFTTFFFFTFTFTTFTFTTFTFFFFTTFTTFFTTTFTFTTTFFTTTTFTTTTFFTTTTTFTFTTTTFTTTFFFTFFFFTFTTTFTTFTFFFFFTFFFTFFTTFFFTFFTFFFFTFFFTFFFTFTTFFFFFTTFTTFTFFFFFTTTTTTTFFFTTTTFTTFFTFTTFFFFFTTFFFFFFFFFFTTFFTTFFFFFFTFTTFTFFTFFFTFFTTTFTTFFFFTFTTTFTTFTFFFTTTFFTTFTTTFFFTTFTFTTTTTFTTFTTFTTTFFFFFFTFTFTTFFTTFTFTTFTFFFTTTFTTTTTFFFFTFTFFTTTFTTTFFTTTFFFFFTTTTFTFFFTFTFFFFFTFFTTTFFTTFFTTTFFFTFTFFTTTTTFFTTTFTTTTFFFFTFFFTFFTTTTTTFTFFFFFTTTTFTTFFTFTTTTFFFTTFTTTFFFFTFTTFTFFTFFTTTFTFFTTTFTTTFTFFFTFFFFTFTFTTTTTFFTFTTTFTFFFTFFTFFFFTFTFFTFFFTFFTTFFFTTFTTFTTFTTFTTTTFTTFTTFTTTTFFTTFFFFFTTFFTFTTTFFFFFTFFTFFFTTTTFTTFFTFFTFTTFFTFFFFFTFTTTFFFFFFFTFFFTTFFFTTFTTTTTFFFTTFFTTFTTFTTTTTFTTFFTFFTTFFTFTTTFFFFTTTTFTFFTFTTTTTTFFFFTTFFFTFFFFFTFFTTTFTFFTFTTTFFFTTFTTFTTFFFFFTFTTTFFFTFFTFFTFFTTTTFTFFFTFFTTTTTFTTTTFFFTFTTTFFTTFFFFTTFFTFTFFFFFTTTTTFFFTTTFTFTFTFTFFFTFFFTTFTTTTFFTFTFTTFTTTTTFTFFFTFTFFFTFFTTTTFFFTTFTFFFTFTFFTTFTFTTTTTFTFTFTFTTTFFTTFTFFTTTTTTFFTTTFFFFFFFTTFTTTFFFTFFFTFTTTTTTTTFTTTFFTTFTFTTTFFTFFTTTTTFTFTTFFTFFFTFTTFFFFFTFFTTTFFTFTTTTTFTTFTFFTFFTFFFTFTTFFTFTTTTTFTTFFTFTFTTFFTTFTTTFTTTFTFFTFTFTFFFTFFFTFFFFFTTTTFFTFFFFTTFFFTTTTFTFFTFTFTTTFFFFTTFTFFTFTFTFTFTTTFTFFFTTTTTTTFTTFFFFTTTTFFFTFTTTFTTTTFTTTFFFTTTTTTFFFTTTTTTFTFTTFTTFFTFFFTTTFFTTTTFTFFFTTFFTTTFFTTTFFTTTFTTTTFFFFFTTFTTTFFFTTFFTFFFFFFFTTFFTTTTTFTTFFTTFFTFFFFTFTFTTFFFTTTFFFFTTFTFTTFFFTTTFTFFTFTFFFTTTFTFTFFTTFFFFTFFFTFFFFTTTTTFTTFTFFTTTFFTFFFTFTTFFFFTFFFTFTFFFTFFFTTTTFFTTTTFFTFFTTFTTFFFFFTTTFTTFFTFFTFFTFTTTTTFFTTTTFFFFTTTFFTFTTTFTFFTTFFFFTFTFTFTTTTTFFTTTTTTFTFTTFTTTFTTTFFTFFTFTFTFFTTFTFFFTFTFFFTFFTFTFTTFFTTTTTFFTFTTFTTFTFFTFTFTFFFTFTFFFFFFFFFFTTFFFTTFTFTFFFTFFFTFFTFFFTTTTTFTFTFFFTFTTFTTTFFTFTFTFTFFTTFTTTFTTFTFFTFFTTFTFFFFFFFFFTTFTTFFTFTTTTFTTTTFTFFTFFFTFTFFFFTTTFTTTTFTTTTTFFFFTTFFFFFTTTFFTTFTTFTFTTFTTTTFTFFTTFTFFFFTFTFFFTTTFTTTTFTTFTFTTFFTTTFTFTFFTTFFTTFTTTFTTTFTTTFFTTFFTTTTTFFFTFFFTTTTTTFTFTFFTTFTTTTFFFFFFFFTFTFFTFFTTFFFFTFTFFTTFTTFTFFTFFTTFFFTFTFTFFFFTFFFFFFTTTFFFTFFTFTFFTFFTTTFTTFFFTFTFTFFTFTFTFFTFFFTFTTTTFTFFTTFFFTTTTTTFFFTFTTFTFFFTTTFFTFTTTTFTTFFTTFFFFFTFTFTTFFFTFTTFTTFTFFFTFFTFTFFTTFFFFTFFFFFTTFFFTFTFFTTFTFTFTFTTTTTTTTFTTFTTTTFFFFFFFTFFTFFFTFFFTTTFFFFFFFFFFFFFFFFTFTTFTFFFFTTFFTFFTFTFFFFTTTFTFTFTFTFFTTFFTFFFTTTTTFTTTTTTTTTTTFTFFFFFFFFTFFTFTFTFTFFFFFTFTFTFTTFFFFFTTFFFTFTFFTTTFTFFFFFTTFFTFFTFFTFFTTTTFTTFTTTFTTTTFTTFFFTFFFTTTTFFFFTTTFTTFTTFTFTTFTTTFFFFTFTTTFFTFTFFFTTTTTFTTFFTFTTTTFFTTTFTTFFTFFFFFFFTFTFFFTFFFTFFTTFFFTFTTTFTFTTTFTTTFTTTFTFFTFFFTTTTTTFTTTFFTFFFTFTTTFTTTFFTTTFFTFFTFFFFTTFTFTTTTFTTFFTFTFTFFFTTTTFTTTTTTFTTFFFFFFFFFTFTFFTTFFFTTTFFFFFTFFFTFTFFFTFTFTFFFFFFFFTTTTFFTTFFFTTTTTTFTFFFFFTFTTFFFTFTFTFFTTTFTTTFTTTTTTFTFTTTFTTFFTFTFFFTTFTFFFFTTFTFFFFFFFTFTFTFTFTFFTFFTFTFFTTTFFTFTFFTFTFFTTTFFFFFTTFTFFTTFFTFFTFFFTTFTFFFTTFFTTTTTTTFTFFFTTTFFTTTTTTTFTTFTTTTTTFTTTTTFTTFTTTTFTTTTFFTTTTFTTTFFFFTTTTTFFFTTFFTFTFTFTFFFFFTTFTFTTFTTTFTTTTTFTTTTTFTTFTTFTTTFTFFTFTTTTTTTFFFFTTTFFTTTFFTTFFFTTTFTFTFFTFFTTTTTFTFFTTFTFTTFTTFTFFFTFFTFFFFTTFFFFFTTTTFFFFTFTFFTTFTFTFFTTTTTTTTFTTTTTFTFTFTTTFTTFTFFTTFFFTTFTFFFFFFTTTFTFTFFTTTTTTFFFFFFFFTFTFFTTFTTTFFFFTFFTTTTTFFTFFTFFTFTTTFTFFTFFTFTTTFFTTFTTFFTFFFFFTFTTFFFTTTFFFTTFTFTTFTFTTTFFTTFTTFTTFFTFFTFTFTFTTTFFTFFFFFTFTFFTTTTFFFFFFFFTTFTFFTTFTT"
	//k := 5113

	res := maxConsecutiveAnswers(answerKey, k)
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMinInMap(m map[string]int) (string, int) {
	if m["T"] < m["F"] {
		return "T", m["T"]
	}
	return "F", m["F"]
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	m := make(map[string]int)
	m["T"] = 0
	m["F"] = 0
	res := 0
	for l, r := 0, 0; r < len(answerKey); r++ {
		val := m[string(answerKey[r])]
		val++
		m[string(answerKey[r])] = val

		_, minVal := findMinInMap(m)
		if minVal > k {
			val = m[string(answerKey[l])]
			val--
			m[string(answerKey[l])] = val

			l++
			continue
		}
		res = max(res, (r-l)+1)
	}
	return res
}
