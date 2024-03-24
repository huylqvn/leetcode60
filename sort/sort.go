package main

import "fmt"

func main() {
	a := []int{8, 7, 4, 3}
	b := mergeSort(a)

	fmt.Println(b)

	a = []int{2, 5, 8, 4}
	b = quickSort(a)
	fmt.Println(b)
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

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	pivot := a[0]
	var left, right []int
	for _, v := range a[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}
