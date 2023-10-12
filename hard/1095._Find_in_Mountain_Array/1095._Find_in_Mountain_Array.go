package main

import "fmt"

type MountainArray struct {
	slice []int
}

func (this *MountainArray) get(index int) int {
	return this.slice[index]
}

func (this *MountainArray) length() int {
	return len(this.slice)
}

func main() {
	//obj := MountainArray{slice: []int{0, 5, 3, 1}} // 3
	//target := 1
	obj := MountainArray{slice: []int{1, 2, 3, 4, 5, 3, 1}}
	target := 3 // 2
	//obj := MountainArray{slice: []int{0, 1, 2, 4, 2, 1}}
	//target := 3 // -1
	fmt.Println(findInMountainArray(target, &obj))
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	length := mountainArr.length()

	l, r := 1, length-1
	for l != r { // find peak index
		midIdx := (r + l) / 2
		if mountainArr.get(midIdx) < mountainArr.get(midIdx+1) {
			l = midIdx + 1
		} else {
			r = midIdx
		}
	}
	peakIdx := l
	l, r = 0, peakIdx

	for l != r { // find on the left side of peak
		midIdx := (r + l) / 2
		if mountainArr.get(midIdx) < target {
			l = midIdx + 1
		} else {
			r = midIdx
		}
	}
	if mountainArr.get(l) == target {
		return l
	}

	l, r = peakIdx+1, length-1
	for l != r { // find on the right side of peak
		midIdx := (r + l) / 2
		if mountainArr.get(midIdx) > target {
			l = midIdx + 1
		} else {
			r = midIdx
		}
	}
	if mountainArr.get(l) == target {
		return l
	}

	return -1
}
