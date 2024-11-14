package bubblesort

// BubbleSort performs bubble sort on a slice of integers
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr) - 1; i++{
		switched := false
		for j := 0; j < len(arr) - i - 1; j++{
			if arr[j] > arr[j + 1] {
				switched = true
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}

		if !switched {
			break
		}
	}
	return arr
}
