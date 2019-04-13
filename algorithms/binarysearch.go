package main

import (
	"fmt"
)

func main() {
	s := []int{1, 4, 7, 8, 9, 12, 25, 82}
	fmt.Println(binarySearch(8, s))
}

func binarySearch(num int, arr []int) bool {

	fmt.Println(arr)

	size := len(arr)

	if size > 1 {
		middle := arr[size/2]
		if middle == num {
			return true
		} else if middle > num {
			return binarySearch(num, arr[:size/2])
		} else {
			return binarySearch(num, arr[size/2:])
		}

	} else {
		if num == arr[0] {
			return true
		} else {
			return false
		}
	}
}
