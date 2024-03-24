package main

import "fmt"

func main() {
	fmt.Println(isValid("((()))"))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	q := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			q = append(q, s[i])
		} else {
			if len(q) == 0 {
				return false
			}
			switch s[i] {
			case ')':
				if q[len(q)-1] == '(' {
					q = q[:len(q)-1]
				} else {
					return false
				}
			case ']':
				if q[len(q)-1] == '[' {
					q = q[:len(q)-1]
				} else {
					return false
				}
			case '}':
				if q[len(q)-1] == '{' {
					q = q[:len(q)-1]
				} else {
					return false
				}
			}
		}
	}
	if len(q) == 0 {
		return true
	}
	return false
}
