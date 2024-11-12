package solution

import (
	"strings"
	"unicode"
)

func trim(s string) string {
	var soft string
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i])) {
			soft += strings.ToLower(string(s[i]))
		}
	}
	return soft
}

func isPalindrome(s string) bool {
	s = trim(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
