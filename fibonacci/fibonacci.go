package fibonacci

func fibonacci(n int) int64 {
	fib := make([]int64, n)
	fib[0] = 1
	fib[1] = 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n-1]
}
