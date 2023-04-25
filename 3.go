package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	letterMap := make(map[rune]bool)
	max := 0

	j := 0
	for i, char := range s {
		if _, ok := letterMap[char]; ok {
			for {
				delete(letterMap, rune(s[j]))
				if rune(s[j]) == char {
					j++
					break
				}
				j++
			}
		}
		if i-j > max {
			max = i - j
		}
		letterMap[char] = true
	}

	return max + 1
}
