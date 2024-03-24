package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	max := nums[0]
	newmax := max

	for i := 1; i < len(nums); i++ {
		if newmax+nums[i] > nums[i] {
			newmax += nums[i]
		} else {
			newmax = nums[i]
		}

		if newmax > max {
			max = newmax
		}
	}
	return max
}
