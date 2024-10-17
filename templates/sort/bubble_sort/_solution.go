package bubblesort

// BubbleSort performs bubble sort on a slice of integers
func BubbleSort(arr []int) []int {
	n := len(arr)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				// Swap elements
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		// Optimization: after each pass, the largest unsorted element is at the end
		n--
	}

	return arr
}
