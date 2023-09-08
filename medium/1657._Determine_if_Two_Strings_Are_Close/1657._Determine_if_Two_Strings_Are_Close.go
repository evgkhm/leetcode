package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	word1 := "cabbba"
	word2 := "abbccc"
	fmt.Println(closeStrings(word1, word2))
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	map1 := make(map[string]int)
	for _, val := range word1 {
		map1[string(val)]++
	}

	map2 := make(map[string]int)
	for _, val := range word2 {
		map2[string(val)]++
	}

	var slice1 []int
	for key, val := range map1 {
		_, ok := map2[key]
		if !ok { // check alphabet is the same
			return false
		}

		slice1 = append(slice1, val)
	}

	var slice2 []int
	for _, val := range map2 {
		slice2 = append(slice2, val)
	}

	sort.SliceStable(slice1, func(i, j int) bool {
		return slice1[i] < slice1[j]
	})
	sort.SliceStable(slice2, func(i, j int) bool {
		return slice2[i] < slice2[j]
	})

	return reflect.DeepEqual(slice1, slice2)
}
