package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("(]"))
}

func isValid(s string) bool {
	parentheses := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			parentheses = append(parentheses, s[i])
		} else if len(parentheses) > 0 {
			properClosingParanthesOfLastPushed := getClosingParanthes(parentheses[len(parentheses)-1])
			if properClosingParanthesOfLastPushed == s[i] {
				parentheses = parentheses[:len(parentheses)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return len(parentheses) == 0
}

func getClosingParanthes(paranthes byte) byte {
	switch paranthes {
	case '(':
		return ')'
	case '{':
		return '}'
	case '[':
		return ']'
	}

	return '_'
}
