package main

import "fmt"

func main() {

	fmt.Println("Hi fibo")
	fmt.Println(fibonacci(13))
	fib := []int{}
	// var fib []int
	fmt.Println(fib)

}

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
