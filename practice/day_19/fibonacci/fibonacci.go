package fibonacci

// GenerateFibonacci generates the Fibonacci sequence up to n terms using dynamic programming
func GenerateFibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	if n == 1 {
		return []int{0}
	}

	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}

// FibonacciTerm calculates the nth term of the Fibonacci sequence
func FibonacciTerm(n int) int {
	// TODO: first term is 0
	if n <= 1 {
		return 0 
	}
	a, b := 0, 1

	for i := 2; i < n; i++{
		a, b = b, a + b
	}
	
	return b
}
