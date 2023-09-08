package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	grid := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	fmt.Println(equalPairs(grid))
}

func equalPairs(grid [][]int) int {
	m := make(map[string]int)

	for _, val := range grid {
		var str strings.Builder
		for _, v := range val {
			s := strconv.Itoa(v)
			str.WriteString(s)
			str.WriteString(",")
		}
		m[str.String()]++
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		var str strings.Builder
		for j := 0; j < len(grid); j++ {
			s := strconv.Itoa(grid[j][i])
			str.WriteString(s)
			str.WriteString(",")
		}

		localRes := 0
		valInMap, ok := m[str.String()]
		if ok {
			if valInMap > 1 {
				localRes += valInMap
			} else {
				localRes = 1
			}

			res += localRes
		}
	}

	return res
}
