package binarysearch

// BinarySearch performs binary search on a sorted slice of integers
func BinarySearch(arr []int, target int) int {
	return bs(&arr, 0, len(arr)-1, target)
}

func bs(arrP *[]int, low, high, target int) int {
	arr := *arrP
	if low > high {
		return -1
	}

	middle := low + (high-low)/2

	if target == arr[middle] {
		return middle
	} else if target < arr[middle] {
		return bs(arrP, low, middle-1, target)
	} else {
		return bs(arrP, middle+1, high, target)
	}
}
