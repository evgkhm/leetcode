package main

import "fmt"

type MyHashMap struct {
	m map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{m: make(map[int]int)}
}

func (this *MyHashMap) Put(key int, value int) {
	this.m[key] = value
}

func (this *MyHashMap) Get(key int) int {
	val, ok := this.m[key]
	if ok {
		return val
	} else {
		return -1
	}
}

func (this *MyHashMap) Remove(key int) {
	delete(this.m, key)
}

func main() {
	obj := Constructor()
	obj.Put(7, 3)
	param2 := obj.Get(7)
	fmt.Println(param2)
	//obj.Remove(2)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
