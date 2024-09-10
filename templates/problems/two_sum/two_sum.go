package twosum

// TwoSum finds two numbers in the array that add up to the target sum
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, found := numMap[complement]; found {
			return []int{j, i}
		}
		numMap[num] = i
	}

	// If no solution is found
	return []int{}
}
