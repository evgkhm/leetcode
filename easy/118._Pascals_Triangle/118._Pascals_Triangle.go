package main

import "fmt"

func main() {
	n := 2
	fmt.Println(generate(n))
}

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	res := make([][]int, 2)
	res[0] = append(res[0], 1)
	res[1] = append(res[1], 1)
	res[1] = append(res[1], 1)

	if numRows == 2 {
		return res
	}

	for n := 2; n < numRows; n++ {
		tmpSlice := make([]int, n+1)
		tmpSlice[0] = 1
		for i := 1; i < n; i++ {
			value := res[n-1][i-1] + res[n-1][i]
			tmpSlice[i] = value
		}
		tmpSlice[n] = 1
		res = append(res, tmpSlice)
	}

	return res
}
