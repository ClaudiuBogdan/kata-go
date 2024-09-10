package linearsearch

// LinearSearch performs a linear search on a slice of integers
func LinearSearch(arr []int, target int) int {
	for i, value := range arr {
		if value == target {
			return i // Target found, return its index
		}
	}
	return -1 // Target not found
}
