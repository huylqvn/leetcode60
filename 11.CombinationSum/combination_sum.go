package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum([]int{2, 5, 8, 4}, 10))
}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	candidates = sort(candidates)
	return combinationSumExe(candidates, target)
}

func combinationSumExe(candidates []int, target int) [][]int {
	result := [][]int{}

	for i, v := range candidates {
		if v == target {
			result = append(result, []int{v})
		}
		if v < target {
			for _, r := range combinationSumExe(candidates[i:], target-v) {
				result = append(result, append([]int{v}, r...))
			}
		} else {
			break
		}
	}

	return result
}

func sort(a []int) []int {
	return mergeSort(a)
}

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	mid := len(a) / 2

	return merge(mergeSort(a[:mid]), mergeSort(a[mid:]))
}

func merge(a, b []int) []int {
	var c []int
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	c = append(c, a[i:]...)
	c = append(c, b[j:]...)
	return c
}
