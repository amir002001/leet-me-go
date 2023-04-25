package main

func characterReplacement(s string, k int) int {
	max := 0
	freqMap := make(map[rune]int)
	maxFreq := 0
	j := 0
	for i, char := range s {
		freqMap[char]++
		if freqMap[char] > maxFreq {
			maxFreq = freqMap[char]
		}
		for (i-j+1)-maxFreq > k {
			freqMap[rune(s[j])]--
			j++
		}
		if max < i-j+1 {
			max = i - j + 1
		}
	}
	return max
}
