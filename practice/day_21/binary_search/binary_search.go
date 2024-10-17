package binarysearch

// BinarySearch performs binary search on a sorted slice of integers
func BinarySearch(arr []int, target int) int {
	return binarySearchAux(&arr, 0, len(arr)-1, target)
}

func binarySearchAux(arrPointer *[]int, low, high, target int) int {
	arr := *arrPointer
	if low > high {
		return -1
	}

	middle := low + (high-low)/2

	if arr[middle] == target {
		return middle
	// TODO: if target is lower that middle value
	} else if target < arr[middle] {
		return binarySearchAux(arrPointer, low, middle-1, target)
	} else {
		return binarySearchAux(arrPointer, middle+1, high, target)
	}
}
