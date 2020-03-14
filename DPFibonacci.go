package main

import "fmt"

func FibNaive(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return FibNaive(n-1) + FibNaive(n-2)
	}
}

var memo [100]int

func FibDPTopDown(n int) int {
	r := 0

	if memo[n] != 0 {
		return memo[n]
	}
	if n == 1 || n == 2 {
		r = 1
	} else {
		r = FibDPTopDown(n-1) + FibDPTopDown(n-2)
	}
	memo[n] = r
	return r
}

func FibDPBottomUp(n int) int {
	r := 0
	var fib = make([]int, n+1) //new([n]int)

	for i := 1; i <= n; i++ {
		if i <= 2 {
			r = 1
		} else {
			r = fib[i-1] + fib[i-2]
		}
		fib[i] = r
	}
	return fib[n]
}

func main() {

	n := 45

	fmt.Print("The result of Function Fibonacci (Naive Version) is: " + fmt.Sprint(FibNaive(n)) + "\n")
	fmt.Print("The result of Function Fibonacci (Dynamic Programing - Memoization (Top Down)) is: " + fmt.Sprint(FibDPTopDown(n)) + "\n")
	fmt.Print("The result of Function Fibonacci (Dynamic Programing - Bottom Up) is: " + fmt.Sprint(FibDPBottomUp(n)) + "\n")

}
