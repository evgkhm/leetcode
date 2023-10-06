package main

import (
	"fmt"
	"strings"
)

func main() {
	//strs := []string{"flower", "flow", "flight"}
	strs := []string{"ab", "a"}
	//strs := []string{""}
	//strs := []string{"flower", "flower", "flower", "flower", "flower"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs[0]) == 0 {
		return ""
	}

	var res strings.Builder

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || i >= len(strs[j-1]) || strs[j][i] != strs[j-1][i] {
				return res.String()
			}
		}
		res.WriteByte(strs[0][i])
	}
	return res.String()
}
