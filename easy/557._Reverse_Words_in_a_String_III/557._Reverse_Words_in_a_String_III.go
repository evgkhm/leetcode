package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest" //"s'teL ekat edoCteeL tsetnoc"
	//s := "God Ding" //"doG gniD"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	if len(s) == 1 {
		return s
	}

	l, r := 0, 0
	idx := 0
	var res strings.Builder
	for idx < len(s) {
		if s[idx] == ' ' {
			for j := r - 1; j >= l; j-- {
				res.WriteByte(s[j])
			}
			res.WriteRune(' ')
			l = r + 1
		} else if idx == len(s)-1 {
			for j := r; j >= l; j-- {
				res.WriteByte(s[j])
			}
			return res.String()
		}

		idx++
		r++
	}
	return res.String()
}
