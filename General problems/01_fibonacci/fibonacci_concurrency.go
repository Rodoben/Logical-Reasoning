package main

import "fmt"

func main() {
	fmt.Println("hi")
	for n := range fibonacciConcurrent(10) {
		fmt.Println(n)
	}

}

func fibonacciConcurrent(c int) <-chan int {
	out := make(chan int)
	go func() {
		a, b := 0, 1
		for i := 0; i < c; i++ {
			out <- a
			a, b = b, a+b
		}
		close(out)
	}()

	return out
}
