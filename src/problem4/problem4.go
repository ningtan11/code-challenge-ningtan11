package main

import "fmt"

// Assumes all input n is positive
// Does not error on negative input but behaviour is undefined

// iterative, time complexity O(n), space complexity O(1)
func sum_to_n_a(n int) int {
	out := 0
	for i := 1; i <= n; i++ {
		out += i
	}
	return out
}

// recursive, time complexity O(n), space complexity O(n)
func sum_to_n_b(n int) int {
	if n <= 1 {
		return n
	}
	return n + sum_to_n_b(n-1)
}

// formula, time complexity O(1), space complexity O(1)
func sum_to_n_c(n int) int {
	return (n * (n + 1)) / 2
}

func main() {
	fmt.Println(sum_to_n_a(0))
	fmt.Println(sum_to_n_b(0))
	fmt.Println(sum_to_n_c(0))
}
