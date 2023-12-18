package main

import "fmt"

func main() {

	fmt.Println("Hi fibo")
	//fmt.Println(fibonacci(13))
	//fib := []int{}
	// var fib []int
	//fmt.Println(fib)

	fibonacciLoop(10)

}

// using a slice
func fibonacci(n int) []int {
	//fib := []int{}
	fib := make([]int, n)
	if n >= 1 {
		fib[0] = 0
	}
	if n >= 2 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

// just using for loop
func fibonacciLoop(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}
