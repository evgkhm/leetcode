package main

import (
	"fmt"
	"strconv"
)

const ()

func main() {
	digits := "234"
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) []string {
	var res []string

	keyboard := make(map[int32]string)
	keyboard[2] = "abc"
	keyboard[3] = "def"
	keyboard[4] = "ghi"
	keyboard[5] = "jkl"
	keyboard[6] = "mno"
	keyboard[7] = "pqrs"
	keyboard[8] = "tuv"
	keyboard[9] = "wxyz"

	createRes(keyboard, &res, digits)

	return res
}

func createRes(keyboard map[int32]string, res *[]string, digits string) {
	if len(digits) == 0 {
		return
	}

	number, _ := strconv.Atoi(string(digits[0]))
	digits = digits[1:]

	lettersInMap := keyboard[int32(number)]

	if *res == nil {
		for _, val := range lettersInMap {
			*res = append(*res, string(val))
		}
	} else {
		var newSlice []string

		for _, lettersFromRes := range *res {
			for _, newLetterFromMap := range lettersInMap {
				var newResultLetter string
				newResultLetter = lettersFromRes + string(newLetterFromMap)
				newSlice = append(newSlice, newResultLetter)
			}
		}
		*res = newSlice
	}

	if len(digits) > 0 {
		createRes(keyboard, res, digits)
	}
}
