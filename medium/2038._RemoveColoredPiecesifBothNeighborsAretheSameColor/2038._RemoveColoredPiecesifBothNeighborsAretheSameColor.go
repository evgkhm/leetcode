package main

import (
	"fmt"
)

func main() {
	//colors := "AAABABB" // true
	colors := "ABBBBBBBAAA"
	fmt.Println(winnerOfGame(colors))
}

func winnerOfGame(colors string) bool {
	countA, countB := 0, 0

	for i := 2; i < len(colors); i++ {
		if colors[i] == 'A' && colors[i-1] == 'A' && colors[i-2] == 'A' {
			countA += 1
		} else if colors[i] == 'B' && colors[i-1] == 'B' && colors[i-2] == 'B' {
			countB += 1
		}
	}

	if countA > countB {
		return true
	}

	return false
}
