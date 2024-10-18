package binarysearch

// BinarySearch performs binary search on a sorted slice of integers
func BinarySearch(arr []int, target int) int {
	return binarySearchAux(&arr, 0, len(arr) - 1, target)
}

func binarySearchAux(arrP *[]int, lower, higher, target int) int {
	arr := *arrP // Dereference arr pointer
	if lower > higher{
		return -1
	}

	middle := lower + (higher - lower)/2
	if arr[middle] == target {
		return middle
	} else if target < arr[middle] {
		return binarySearchAux(arrP, lower, middle - 1, target)
	} else {
		return binarySearchAux(arrP, middle + 1, higher, target)
	}
}
