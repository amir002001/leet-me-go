package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var builder strings.Builder
	for _, r := range s {
		if isAlphanum(r) {
			builder.WriteRune(unicode.ToLower(r))
		}
	}
	s = builder.String()
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func isAlphanum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
