package _979__Find_Greatest_Common_Divisor_of_Array

func findGCD(nums []int) int {
	min, max := 1000, 1
	for _, val := range nums {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return nod(min, max)
}

// nod
func nod(a int, b int) int {
	for a != 0 && b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	return a + b
}
