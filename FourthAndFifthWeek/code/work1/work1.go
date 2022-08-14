package main

import "fmt"

func isValid(s string) bool {
	var leftBrackets = make([]byte, 0, len(s))
	leftBrackets = append(leftBrackets, '#')
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '{', '[':
			leftBrackets = append(leftBrackets, s[i])
		case ')':
			if leftBrackets[len(leftBrackets)-1] == '(' {
				leftBrackets = leftBrackets[:len(leftBrackets)-1]
				continue
			} else {
				return false
			}
		case '}':
			if leftBrackets[len(leftBrackets)-1] == '{' {
				leftBrackets = leftBrackets[:len(leftBrackets)-1]
				continue
			} else {
				return false
			}
		case ']':
			if leftBrackets[len(leftBrackets)-1] == '[' {
				leftBrackets = leftBrackets[:len(leftBrackets)-1]
				continue
			} else {
				return false
			}
		}
	}
	if len(leftBrackets) > 1 {
		return false
	}
	return true
}

func main() {
	fmt.Println("\"[\":\t", isValid("["))
	fmt.Println("\"(]\":\t", isValid("(]"))
	fmt.Println("\"()[]{}\":\t", isValid("()[]{}"))
	fmt.Println("\"([)]\":\t", isValid("([)]"))
	fmt.Println("\"{[]}\":\t", isValid("([)]"))
}
