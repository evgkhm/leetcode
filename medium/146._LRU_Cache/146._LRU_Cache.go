package main

type LRUCache struct {
	m             map[int]int
	lastUsedSlice []int
	cap           int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{m: make(map[int]int, capacity), lastUsedSlice: []int{}, cap: capacity}
	return lru
}

func (l *LRUCache) Get(key int) int {
	res := -1

	value, ok := l.m[key]
	if ok {
		res = value

		for i := 0; i < len(l.lastUsedSlice); i++ {
			if l.lastUsedSlice[i] == key {
				l.lastUsedSlice = append(l.lastUsedSlice[0:i], l.lastUsedSlice[i+1:]...)
			}
		}

		l.lastUsedSlice = append(l.lastUsedSlice, key)
	}
	return res
}

func (l *LRUCache) Put(key int, value int) {
	l.lastUsedSlice = append(l.lastUsedSlice, key)

	if len(l.m) < l.cap {
		l.m[key] = value
	} else if l.cap == 1 { // если вместимость мапы только 1 то просто перезапись
		for k := range l.m {
			delete(l.m, k) //убрать старый key value
		}
		l.m[key] = value // перезапись key value на новый
	} else {
		if len(l.lastUsedSlice) > 0 {
			_, ok := l.m[key] // обновляем значение value если такой ключ уже есть
			if ok {
				l.m[key] = value
			} else {
				// если такого ключа нет и длина len уже закончилась то нужно
				// найти неиспользованный ключ и заменить его на новый key value
				// и убираем из lastUsedSlice первый элемент, т.к. он самый неиспользованный LRU
				for k, _ := range l.m {
					if k == l.lastUsedSlice[0] {
						delete(l.m, k)                        // удаляем старый ключ
						l.m[key] = value                      // записываем новый
						l.lastUsedSlice = l.lastUsedSlice[1:] //убираем использованный
						break
					}
				}
			}
		}
	}

	if len(l.lastUsedSlice) > len(l.m) {
		// сюда попадем если нужно убрать элемент из lastUsedSlice где-то внутри слайса
		for i := 0; i < len(l.lastUsedSlice); i++ {
			if l.lastUsedSlice[i] == key {
				l.lastUsedSlice = append(l.lastUsedSlice[0:i], l.lastUsedSlice[i+1:]...)
				break
			}
		}
	}
}

func main() {
	//TEST1
	//obj := Constructor(2)
	//obj.Put(1, 1)
	//obj.Put(2, 2)
	//obj.Get(1)
	//obj.Put(3, 3)
	//obj.Get(2)
	//obj.Put(4, 4)
	//obj.Get(1)
	//obj.Get(3)
	//obj.Get(4)

	//TEST9
	//obj := Constructor(1)
	//obj.Put(2, 1)
	//obj.Get(2)
	//obj.Put(3, 2)
	//obj.Get(2)
	//obj.Get(3)

	////TEST11
	//obj := Constructor(2)
	//obj.Put(2, 1)
	//obj.Put(2, 2)
	//obj.Get(2)
	//obj.Put(1, 1)
	//obj.Put(4, 1)
	//obj.Get(2)

	////TEST13
	//obj := Constructor(2)
	//obj.Get(2)
	//obj.Put(2, 6)
	//obj.Get(1)
	//obj.Put(1, 5)
	//obj.Put(1, 2)
	//obj.Get(1)
	//obj.Get(2)

	//TEST
	//obj := Constructor(3)
	//obj.Put(1, 1)
	//obj.Put(2, 2)
	//obj.Put(3, 3)
	//obj.Put(4, 4)
	//obj.Get(4)
	//obj.Get(3)
	//obj.Get(2)
	//obj.Get(1)
	//obj.Put(5, 5)
	//obj.Get(1)
	//obj.Get(2)
	//obj.Get(3)
	//obj.Get(4)
	//obj.Get(5)

	//TEST14
	//obj := Constructor(2)
	//obj.Put(2, 1)
	//obj.Put(1, 1)
	//obj.Put(2, 3)
	//obj.Put(4, 1)
	//obj.Get(1)
	//obj.Get(2)

	//TEST15
	obj := Constructor(10)
	obj.Put(10, 13)
	obj.Put(3, 17)
	obj.Put(6, 11)
	obj.Put(10, 5)
	obj.Put(9, 10)
	obj.Get(13)
	obj.Put(2, 19)
	obj.Get(2)
	obj.Get(3)
	obj.Put(5, 25)
	obj.Get(8)
	obj.Put(9, 22)
	obj.Put(5, 5)
	obj.Put(1, 30)
	obj.Get(11)
	obj.Put(9, 12)
	obj.Get(7)
	obj.Get(5)
	obj.Get(8)
	obj.Get(9)
	obj.Put(4, 30)
	obj.Put(9, 3)
	obj.Get(9)
	obj.Get(10)
	obj.Get(10)
	obj.Put(6, 14)
	obj.Put(3, 11)
	obj.Get(3)
	obj.Put(10, 11)
	obj.Get(8)
	obj.Put(2, 14)
	obj.Get(1)
	obj.Get(5)
	obj.Get(4)
	obj.Put(11, 4)
	obj.Put(12, 24)
	obj.Put(5, 18)
	obj.Get(13)
	obj.Put(7, 23)
	obj.Get(8)
	obj.Get(12)
	obj.Put(3, 27)
	obj.Put(2, 12)
	obj.Get(5)
	obj.Put(2, 9)
	obj.Put(13, 4)
	obj.Put(8, 18)
	obj.Put(1, 7)
	obj.Get(6)
	obj.Put(9, 29)
	obj.Put(8, 21)
	obj.Get(5)
	obj.Put(6, 30)
	obj.Put(1, 12)
	obj.Get(10)
}
