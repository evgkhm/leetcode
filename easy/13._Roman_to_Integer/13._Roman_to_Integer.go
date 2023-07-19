package _3__Roman_to_Integer

func romanToInt(s string) int {
	res := 0
	prev := ""

	for _, val := range s {
		switch string(val) {
		case "M":
			if prev == "C" {
				res += 800
			} else {
				res += 1000
			}

		case "D":
			if prev == "C" {
				res += 300
			} else {
				res += 500
			}

		case "C":
			if prev == "X" {
				res += 80
			} else {
				res += 100
			}

		case "L":
			if prev == "X" {
				res += 30
			} else {
				res += 50
			}

		case "X":
			if prev == "I" {
				res += 8
			} else {
				res += 10
			}

		case "V":
			if prev == "I" {
				res += 3
			} else {
				res += 5
			}

		case "I":
			res += 1
		}
		prev = string(val)
	}
	return res
}
