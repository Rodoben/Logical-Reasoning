package main

import "fmt"

func main() {
	xi := []int{3, 2, 1, 5, 1, 2, 4, 5}
	fmt.Println(BubbleSort(xi))
}

func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
