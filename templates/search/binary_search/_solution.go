package binarysearch

// BinarySearch performs binary search on a sorted slice of integers
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // Avoid potential integer overflow

		if arr[mid] == target {
			return mid // Target found
		} else if arr[mid] < target {
			left = mid + 1 // Target is in the right half
		} else {
			right = mid - 1 // Target is in the left half
		}
	}

	return -1 // Target not found
}
