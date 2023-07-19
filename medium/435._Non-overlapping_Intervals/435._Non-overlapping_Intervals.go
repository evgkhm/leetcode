package main

import (
	"fmt"
	"sort"
)

func main() {
	//intervals := [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	//intervals := [][]int{{1, 2}, {1, 2}, {1, 2}}
	//intervals := [][]int{{1, 2}, {2, 3}}
	res := eraseOverlapIntervals(intervals)
	fmt.Println(res)
}

func eraseOverlapIntervals(intervals [][]int) int {
	res := 0

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	k := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		x := intervals[i][0]
		y := intervals[i][1]
		if x < k {
			res++
		} else {
			k = y
		}
	}

	return res
}
