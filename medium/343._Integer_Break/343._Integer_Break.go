package main

import "fmt"

func main() {
	n := 10

	fmt.Println(integerBreak(n))
}

func integerBreak(n int) int {
	if n == 2 || n == 3 {
		return n - 1
	}

	res1 := divide(n, 3)
	res2 := divide(n, 2)
	if res1 > res2 {
		return res1
	}
	return res2
}

func divide(n int, factor int) int {
	res := 1

	var values []int
	for n > 0 {
		if n-factor >= 0 {
			values = append(values, factor)
			n -= factor
		} else {
			if n == 1 {
				values[0] += n
				n -= n
			} else if n == 2 {
				values = append(values, n)
				n -= n
			}
		}
	}

	for _, val := range values {
		res *= val
	}

	return res
}
