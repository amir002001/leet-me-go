package main

import "strings"

func isValid(s string) bool {
	openings := "[({"
	closings := "])}"
	stack := []rune{}

	for _, char := range s {
		if strings.ContainsRune(openings, char) {
			stack = append(stack, char)
		} else if strings.ContainsRune(closings, char) {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			lastIdx := strings.IndexRune(openings, last)
			charIdx := strings.IndexRune(closings, char)
			if charIdx != lastIdx {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
