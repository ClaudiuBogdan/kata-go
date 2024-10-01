package quicksort

import "math/rand"

// QuickSort performs quick sort on a slice of integers
func QuickSort(arr []int) []int {
	quickSortHelper(&arr, 0, len(arr)-1)
	return arr
}

func quickSortHelper(arr *[]int, low int, high int) {
	if low >= high {
		return
	}

	pivotIndex := partition(arr, low, high)

	quickSortHelper(arr, low, pivotIndex-1)
	quickSortHelper(arr, pivotIndex+1, high)
}

func partition(arrPointer *[]int, low int, high int) int {
	arr := *arrPointer
	// Get a random number and swap it with the high
	// TODO: practice random number generator
	pivotIndex := rand.Intn(high-low+1) + low
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

	pivot := arr[high]

	index := low - 1

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			index++
			arr[i], arr[index] = arr[index], arr[i]
		}
	}

	index++

	arr[index], arr[high] = arr[high], arr[index]

	return index
}
