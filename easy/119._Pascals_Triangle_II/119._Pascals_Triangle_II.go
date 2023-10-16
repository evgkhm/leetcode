package main

import "fmt"

func main() {
	n := 2
	fmt.Println(getRow(n))
}

func getRow(numRows int) []int {
	if numRows == 0 {
		return []int{1}
	}
	if numRows == 1 {
		return []int{1, 1}
	}

	res := make([]int, 3)
	res[0] = 1
	res[1] = 2
	res[2] = 1

	for n := 2; n < numRows; n++ {
		tmpSlice := make([]int, len(res)+1)
		tmpSlice[0] = 1
		for i := 1; i <= n; i++ {
			value := res[i-1] + res[i]
			tmpSlice[i] = value
		}
		tmpSlice[n+1] = 1
		res = tmpSlice
	}

	return res
}
