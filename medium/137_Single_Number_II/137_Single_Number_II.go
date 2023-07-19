package main

import "fmt"

func main() {
	slice := make([]int, 7)
	slice[0] = 0
	slice[1] = 1
	slice[2] = 0
	slice[3] = 1
	slice[4] = 0
	slice[5] = 1
	slice[5] = 99
	res := singleNumber(slice)
	fmt.Println(res)
}

func singleNumber(nums []int) int {
	m := make(map[int]int)
	counter := 0

	for _, num := range nums {
		_, ok := m[num]
		if ok {
			counter = m[num]
			counter++
			m[num] = counter
		} else {
			m[num] = 1
		}
	}

	res := 0
	for key, val := range m {
		if val == 1 {
			res = key
			return res
		}
	}

	return res
}
