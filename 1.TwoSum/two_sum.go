package main

import "fmt"

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]
func main() {
	fmt.Println(twoSum1([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum2([]int{3, 2, 4}, 6))
}

// 1. Brute Force
func twoSum1(nums []int, target int) []int {
	for i, v := range nums {
		for j, k := range nums[i+1:] {
			if v+k == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return []int{}
}

// 2. Hash Table
func twoSum2(nums []int, target int) []int {
	tempMap := make(map[int]int)
	for i, v := range nums {
		if j, ok := tempMap[target-v]; ok {
			return []int{j, i}
		}
		tempMap[v] = i
	}
	return []int{}
}
