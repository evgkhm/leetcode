package main

import (
	"fmt"
)

func main() {
	//str1, str2 := "LEET", "CODE" // ""
	//str1, str2 := "ABCDEF", "ABC" // ""
	//str1, str2 := "ABCABC", "ABC" // ABC
	str1, str2 := "ABABAB", "ABAB" // AB
	//str1, str2 := "LEET", "CODE" // ""
	//str1, str2 := "AAAAAAAAA", "AAACCC" // ""
	res := gcdOfStrings(str1, str2)
	fmt.Println(res)
}

func gcdOfStrings(str1 string, str2 string) string {
	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}

	if str1 == str2 {
		return str1
	}

	if str1[:len(str2)] != str2 {
		return ""
	}

	return gcdOfStrings(str1[len(str2):], str2)
}
