package tasks

import "fmt"

func binarySearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == key {
			return mid
		} else if arr[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func Task17() {
	arr := []int{1, 2, 7, 7, 9, 10}

	ind := binarySearch(arr, 7)

	fmt.Print(ind)
}
