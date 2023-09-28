package main

import (
	"fmt"
	"sort"
)

type Soldier struct {
	index int
	value int
}

func main() {
	mat := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}
	k := 3
	fmt.Println(kWeakestRows(mat, k))
}

func kWeakestRows(mat [][]int, k int) []int {
	var soldiers []Soldier

	for idx, val := range mat {
		soldierCount := countSoldiers(val)
		soldier := Soldier{index: idx, value: soldierCount}
		soldiers = append(soldiers, soldier)
	}

	sort.SliceStable(soldiers, func(i, j int) bool {
		return soldiers[i].value < soldiers[j].value
	})

	var res []int
	for i := 0; i < k; i++ {
		res = append(res, soldiers[i].index)
	}

	return res
}

func countSoldiers(soldiers []int) int {
	start, end := 0, len(soldiers)-1
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		if soldiers[mid] == 0 {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}
