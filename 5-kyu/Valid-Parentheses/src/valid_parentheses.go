package src

import (
	"strings"
)

//第一种
func ValidParentheses(parens string) bool {
	// Your code goes here
	for strings.Contains(parens, "()") {
		parens = strings.Replace(parens, "()", "", -1)
	}
	return len(parens) == 0
}

//第二种
func ValidParentheses2(parens string) bool {
	// Your code goes here
	count := 0
	for _, c := range parens {
		if c == '(' {
			count++
		} else if c == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}
