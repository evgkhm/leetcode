package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//s := "3[a]2[bc]" //aaabcbc
	//s := "3[a2[c]]" //"accaccacc"
	//s := "10[leetcode]"
	s := "2[abc]3[cd]ef"
	res := decodeString(s)
	fmt.Println(res)
}

func decodeString(s string) string {
	res := ""
	factor := 0
	var prevResStringSlice []string
	var sliceInt []int

	factorStr := ""
	for _, val := range s {
		if unicode.IsNumber(val) == true {
			factorStr += string(val)
			factor, _ = strconv.Atoi(factorStr)
		} else if val == '[' {
			prevResStringSlice = append(prevResStringSlice, res)
			sliceInt = append(sliceInt, factor)
			res = ""
			factorStr = ""
			factor = 0
		} else if unicode.IsLetter(val) == true {
			res += string(val)
		} else if val == ']' {
			idx := len(sliceInt) - 1 //index for pop
			preFactor := sliceInt[idx]
			sliceInt = sliceInt[:idx] //pop

			idx = len(prevResStringSlice) - 1
			preResString := prevResStringSlice[idx]
			prevResStringSlice = prevResStringSlice[:idx] //pop

			res = preResString + strings.Repeat(res, preFactor)
		}
	}

	return res
}
