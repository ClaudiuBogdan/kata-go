package insertionsort

// InsertionSort performs insertion sort on a slice of integers
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Move elements of arr[0..i-1] that are greater than key
		// to one position ahead of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}

	return arr
}
