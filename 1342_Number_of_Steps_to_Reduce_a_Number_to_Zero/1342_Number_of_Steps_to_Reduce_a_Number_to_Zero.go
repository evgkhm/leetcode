package main

import "fmt"

func numberOfSteps(num int) int {
	res := 0
	for num > 0 {
		if num&1 == 0 { //even
			num >>= 1
		} else { //odd
			num -= 1
		}
		res++
	}
	return res
}

func main() {
	input := 14
	res := numberOfSteps(input)
	fmt.Println(res)
}
