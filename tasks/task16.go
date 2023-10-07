package tasks

import "fmt"

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++

			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)

		quickSort(arr, low, pivotIndex-1)

		quickSort(arr, pivotIndex+1, high)
	}
}

func Task16() {
	arr := []int{6, 6, 2, 1, 5, 3, 7, 4}

	quickSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}
