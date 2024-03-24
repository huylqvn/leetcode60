package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))

}

func lengthOfLongestSubstring(s string) int {
	max := 1
	i := 0
	count := 0
	seen := make(map[rune]int)
	for i < len(s) {
		if v, ok := seen[rune(s[i])]; ok {
			count = 0
			i = v + 1
			seen = make(map[rune]int)
			continue
		} else {
			count++
			seen[rune(s[i])] = i
			i++
		}

		if count > max {
			max = count
		}
	}
	return max
}
