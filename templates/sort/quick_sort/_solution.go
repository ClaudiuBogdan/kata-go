package quicksort

import (
	"math/rand"
)

// QuickSort performs quick sort on a slice of integers
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	quickSortHelper(arr, 0, len(arr)-1)
	return arr
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSortHelper(arr, low, pivotIndex-1)
		quickSortHelper(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// Choose a random pivot to avoid worst-case scenarios for already sorted arrays
	pivotIndex := rand.Intn(high-low+1) + low
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]
	pivot := arr[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
