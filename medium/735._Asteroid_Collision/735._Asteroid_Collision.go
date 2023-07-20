package main

import "fmt"

func main() {
	//asteroids := []int{-2, 2, -1, -2} // -2
	//asteroids := []int{5, 10, -5} // 5 10
	//asteroids := []int{10, 2, -5} // 10
	//asteroids := []int{8, -8} // []
	//asteroids := []int{-2, -2, -2, -2} //-2 -2 -2 -2
	//asteroids := []int{-2, -2, -2, 1} //-2 -2 -2 1
	//asteroids := []int{-2, -2, 1, -2} //-2 -2 -2
	//asteroids := []int{-2, 1, 1, -2} // -2,-2
	//asteroids := []int{1, -2, 1, -2} // -2,-2
	asteroids := []int{1, -1, -2, -2}
	res := asteroidCollision(asteroids)
	fmt.Println(res)
}

func asteroidCollision(asteroids []int) []int {
	var res []int

	res = append(res, asteroids[0])

	for i := 1; i < len(asteroids); i++ {
		n := len(res) - 1
		if n < 0 { // res пустой
			res = append(res, asteroids[i])
			continue
		}

		if asteroids[i] < 0 && res[n] > 0 { // если число справа отрицательное а слева положительное
			if asteroids[i]+res[n] < 0 {
				res = res[:len(res)-1]          // удаляем положительное
				res = append(res, asteroids[i]) // добавляем отрицательное
				for n > 0 {                     // надо  проверить, что все остальные элементы в res не уничтожит этот астероид
					if res[n] < 0 && res[n-1] > 0 { // если число справа отрицательное а слева положительное
						if res[n]+res[n-1] < 0 { // негативный астероид уничтожит текущий положительный астероид и займет его место
							res = res[:len(res)-1]
							res = res[:len(res)-1]
							res = append(res, asteroids[i])
						} else if res[n]+res[n-1] > 0 { // положительный астероид уничтожит отрицательный
							res = res[:len(res)-1]
						} else if res[n]+res[n-1] == 0 { // они равны, взаимоуничтожатся
							res = res[:len(res)-1]
							res = res[:len(res)-1]
							break
						}
					} else {
						break
					}
					n--
				}
			} else if res[n] > 0 && asteroids[i] < 0 && asteroids[i]+res[n] == 0 { // летят навстречу друг друга и равны
				if len(res) > 0 {
					res = res[:len(res)-1]
				}
			}
			continue
		}
		res = append(res, asteroids[i]) // если все норм просто добавляем астероид
	}

	return res
}
