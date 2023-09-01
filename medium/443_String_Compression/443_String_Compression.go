package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	chars := []byte{'a', 'b', 'c'}
	//chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	//chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	res := compress(chars)
	fmt.Println(res)
}

func compress(chars []byte) int {
	l, r := 0, 0
	currChar := chars[0]
	var resStr strings.Builder
	length := 0
	for r < len(chars) {
		length = r - l

		if chars[r] != currChar {
			if length > 1 {
				resStr.WriteByte(currChar)
				resStr.WriteString(strconv.Itoa(length))
			} else {
				resStr.WriteByte(currChar)
			}

			currChar = chars[r]
			l = r
		}

		if r == len(chars)-1 { // last char
			length = r - l + 1
			if length > 1 {
				resStr.WriteByte(currChar)
				resStr.WriteString(strconv.Itoa(length))
			} else {
				resStr.WriteByte(currChar)
			}
		}
		r++
	}

	for i := 0; i < len(resStr.String()); i++ {
		chars[i] = resStr.String()[i]
	}

	return resStr.Len()
}
