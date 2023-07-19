package main

import "fmt"

func main() {
	//s := "ABAB"
	//k := 2

	//s := "AABABBA"
	//k := 1

	//s := "AABA"
	//k := 0

	//s := "ABBB"
	//k := 2

	//s := "ABCDE"
	//k := 1

	s := "BAAAB"
	k := 2

	//s := "AAAA"
	//k := 2

	//s := "KJRGKSKKOKLPADMAGODEDNEBMJMKGAPNLSAKADRLHHDRMJTMFBSIQGHENKABISHQNRFJKEPMFIPNDNEOBRJEHFLIHMDLMCIHLHQN"
	//k := 5

	res := characterReplacement(s, k)
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxFrequency(m map[uint8]int) int {
	var maxFreq int
	for _, val := range m {
		if val > maxFreq {
			maxFreq = val
		}
	}
	return maxFreq
}

func characterReplacement(s string, k int) int {
	res := 0

	charMap := make(map[uint8]int)

	for l, r := 0, 0; r < len(s); r++ {
		charMap[s[r]]++

		maxFreq := maxFrequency(charMap)
		windowsSize := r - l + 1

		if windowsSize-maxFreq > k {
			charMap[s[l]]--
			l++
			continue
		}
		res = max(res, (r-l)+1)
	}
	return res
}
