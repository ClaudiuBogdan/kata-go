package fibonacci

// GenerateFibonacci generates the Fibonacci sequence up to n terms using dynamic programming
func GenerateFibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{0, 1}
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp
}

// FibonacciTerm calculates the nth term of the Fibonacci sequence
func FibonacciTerm(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	prev, curr := 0, 1
	for i := 3; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}
