package main

import (
	"fmt"
)

func main() {
	s := "leet**cod*e"
	fmt.Println(removeStars(s))
}

func removeStars(s string) string {
	var res []rune

	for _, val := range s {
		if val != '*' {
			res = append(res, val)
		} else {
			res = res[:len(res)-1]
		}
	}

	return string(res)
}

// second submission using strings.Builder

//func removeStars(s string) string {
//	var res strings.Builder
//
//	for _, val := range s {
//		if val != '*' {
//			res.WriteRune(val)
//		} else {
//			tmp := res.String()
//			tmp = tmp[:len(tmp)-1]
//			res.Reset()
//			res.WriteString(tmp)
//		}
//	}
//
//	return res.String()
//}
