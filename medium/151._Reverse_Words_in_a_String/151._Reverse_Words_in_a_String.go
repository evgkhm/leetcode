package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "the sky is blue"
	res := reverseWords(s)
	fmt.Println(res)
}

func reverseWords(s string) string {
	var tmpStrBuilder strings.Builder
	var resStrBuilder []strings.Builder
	var res string

	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) == true || unicode.IsDigit(rune(s[i])) == true {
			tmpStrBuilder.WriteByte(s[i])
		} else if tmpStrBuilder.Len() > 0 {
			resStrBuilder = append(resStrBuilder, tmpStrBuilder)
			tmpStrBuilder.Reset()
		}
	}

	// add last word
	if tmpStrBuilder.Len() > 0 {
		resStrBuilder = append(resStrBuilder, tmpStrBuilder)
		tmpStrBuilder.Reset()
	}

	// add all strings and spaces
	for i := len(resStrBuilder) - 1; i >= 0; i-- {
		tmpStrBuilder.WriteString(resStrBuilder[i].String())
		tmpStrBuilder.WriteString(" ")
	}

	// delete last space from result
	res = tmpStrBuilder.String()
	if res[len(res)-1] == ' ' {
		res = res[:len(res)-1]
	}

	return res
}
