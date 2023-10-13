package main

import (
	"fmt"
)

func main() {
	//cost := []int{10, 15, 20} //15
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	dp := 0
	dp1 := 0
	dp2 := 0
	for i := 2; i <= len(cost); i++ {
		oneStep := dp1 + cost[i-1]
		twoStep := dp2 + cost[i-2]
		dp = min(oneStep, twoStep)
		dp2 = dp1
		dp1 = dp
	}
	return dp1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
