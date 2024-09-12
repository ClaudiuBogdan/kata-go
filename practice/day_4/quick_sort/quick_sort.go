package quicksort

// QuickSort performs quick sort on a slice of integers
func QuickSort(arr []int) []int {
	return quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, lo int, hi int) []int {
	if lo >= hi {
		return arr
	}
	partitionIdx := partition(arr, lo, hi)

	quickSortHelper(arr, lo, partitionIdx-1)
	quickSortHelper(arr, partitionIdx+1, hi)

	return arr
}

func partition(arr []int, lo int, hi int) int {
	partitionVal := arr[hi]

	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] < partitionVal {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}
	idx++

	arr[idx], arr[hi] = arr[hi], arr[idx]

	return idx
}
