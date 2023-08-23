package main

import (
	"fmt"
)

func main() {
	gain := []int{44, 32, -9, 52, 23, -50, 50, 33, -84, 47, -14, 84, 36, -62, 37, 81, -36, -85, -39, 67, -63, 64, -47, 95, 91, -40, 65, 67, 92, -28, 97, 100, 81}
	res := largestAltitude(gain)
	fmt.Println(res)
}

func largestAltitude(gain []int) int {
	pre := make([]int, len(gain)+1)
	maxHigh := pre[0]

	for i := 0; i < len(gain); i++ {
		pre[i+1] = pre[i] + gain[i]

		if pre[i+1] > maxHigh {
			maxHigh = pre[i+1]
		}
	}

	return maxHigh
}
