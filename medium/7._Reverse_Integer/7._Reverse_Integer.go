package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	x := -901000
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	stringX := strconv.Itoa(x)
	var reversString strings.Builder

	for i := len(stringX) - 1; i >= 0; i-- {
		if (reversString.Len() == 0 && stringX[i] == '0') || stringX[i] == '-' {
			continue
		}
		reversString.WriteByte(stringX[i])
	}

	res, _ := strconv.Atoi(reversString.String())

	if stringX[0] == '-' {
		res = ^res + 1
	}

	if res < -2147483648 || res > 2147483647 {
		return 0
	}

	return res
}
