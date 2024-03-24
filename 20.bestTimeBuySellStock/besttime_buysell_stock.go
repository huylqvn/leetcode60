package main

import "fmt"

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	max := 0
	min := prices[0]

	for _, v := range prices {
		if v < min {
			min = v
		}
		if v-min > max {
			max = v - min
		}
	}
	return max
}
