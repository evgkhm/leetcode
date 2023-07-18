package _05__Can_Place_Flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	l := 0
	counter := 1
	for l < len(flowerbed) {
		if flowerbed[l] == 0 {
			counter++
			if counter == 3 {
				n--
				counter = 1
			}
		} else {
			counter = 0
		}
		l++
	}

	if counter == 2 { //border without plant
		n--
	}

	return n <= 0
}
