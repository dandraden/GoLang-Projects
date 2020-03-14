package main

import "fmt"

func FibNaive(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return FibNaive(n-1) + FibNaive(n-2)
	}
}

func main() {

	fmt.Print("The result of Function Fibonacci (Naive Version) is: " + fmt.Sprint(FibNaive(10)))

}
