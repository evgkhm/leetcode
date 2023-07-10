package main

import (
	"fmt"
	"sort"
)

func main() {
	//weights := make([]int, 4)
	//weights[0] = 1
	//weights[1] = 3
	//weights[2] = 5
	//weights[3] = 1
	//k := 2

	weights := make([]int, 6)
	weights[0] = 3
	weights[1] = 1
	weights[2] = 2
	weights[3] = 6
	weights[4] = 8
	weights[5] = 4
	k := 3

	//weights := make([]int, 2)
	//weights[0] = 1
	//weights[1] = 3
	//k := 2

	res := putMarbles(weights, k)
	fmt.Println(res)
}

func putMarbles(weights []int, k int) int64 {
	if len(weights) <= 2 {
		return 0
	}

	var pairs []int
	for i := 0; i < len(weights)-1; i++ {
		pairs = append(pairs, weights[i]+weights[i+1])
	}

	sort.Ints(pairs)

	res := 0
	for i := 0; i < k-1; i++ {
		max := pairs[len(pairs)-1-i]
		min := pairs[i]
		res += max - min
	}

	return int64(res)
}
