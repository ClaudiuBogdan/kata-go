package quicksort

// QuickSort performs quick sort on a slice of integers
func QuickSort(arr []int) []int {
	quickSort(&arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr *[]int, low int, high int) {
	if low >= high {
		return
	}
	partitionIdx := partition(arr, low, high)
	quickSort(arr, low, partitionIdx-1)
	quickSort(arr, partitionIdx+1, high)
}

func partition(arrPointer *[]int, low int, high int) int {
	// We get the high as the partition index
	arr := *arrPointer
	partitionValue := arr[high]

	idx := low - 1

	for i := low; i < high; i++ {
		if arr[i] <= partitionValue {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}
	idx++
	arr[idx], arr[high] = arr[high], arr[idx]

	return idx
}
