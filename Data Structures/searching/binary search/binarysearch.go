package main

import "fmt"

func main() {

	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2))
	fmt.Println(binarySearchIndex([]int{1, 2, 3, 4, 5, 6, 7, 8}, 27))

}

func binarySearch(xi []int, target int) bool {
	firstindex := 0
	lastindex := len(xi) - 1

	for firstindex <= lastindex {
		midvalue := (firstindex + lastindex) / 2
		if target == xi[midvalue] {
			return true
		}
		if target > xi[midvalue] {
			firstindex = midvalue + 1
		} else {
			lastindex = midvalue - 1
		}
	}

	return false
}

func binarySearchIndex(xi []int, target int) int {
	low, high := 0, len(xi)-1

	for low <= high {
		mid := (low + high) / 2
		if target == xi[mid] {
			return mid
		}

		if target > xi[mid] {
			low = mid + 1
		} else if target < xi[mid] {
			high = mid - 1
		}
	}
	return -1
}
