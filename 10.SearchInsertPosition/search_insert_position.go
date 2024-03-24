package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 4, 5}, 4))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if nums[0] >= target {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	i := len(nums) / 2
	for i >= 0 || i < len(nums) {
		if nums[i] == target {
			return i
		} else if nums[i] > target {
			i--
			if nums[i] < target {
				return i + 1
			}
		} else {
			i++
			if nums[i] > target {
				return i
			}
		}
	}
	return 0
}
