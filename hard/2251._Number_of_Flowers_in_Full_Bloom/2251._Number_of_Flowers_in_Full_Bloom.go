package main

import (
	"fmt"
	"sort"
)

func main() {
	flowers := [][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}
	people := []int{2, 3, 7, 11} //[1 2 2 2]

	//flowers := [][]int{{1, 10}, {3, 3}}
	//people := []int{3, 3, 2}

	fmt.Println(fullBloomFlowers(flowers, people))
}

func fullBloomFlowers(flowers [][]int, people []int) []int {
	var startBloom []int
	var endBloom []int

	for i := 0; i < len(flowers); i++ {
		startBloom = append(startBloom, flowers[i][0])
	}

	for i := 0; i < len(flowers); i++ {
		endBloom = append(endBloom, flowers[i][1]+1) // +1 because flowers stop blooming at end+1, not at end
	}

	sort.Slice(startBloom, func(i, j int) bool {
		return startBloom[i] < startBloom[j]
	})
	sort.Slice(endBloom, func(i, j int) bool {
		return endBloom[i] < endBloom[j]
	})

	var res []int
	for _, person := range people {
		countBlooming := UpperBound(startBloom, person)

		countStoppedBlooming := UpperBound(endBloom, person)

		res = append(res, countBlooming-countStoppedBlooming)
	}

	return res
}

func UpperBound(slice []int, target int) int {
	l, r := 0, len(slice)-1

	for l <= r {
		mid := (r + l) / 2

		if slice[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}
