package main

func isAnagram(s string, t string) bool {
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	for _, char := range s {
		sMap[char]++
	}
	for _, char := range t {
		tMap[char]++
	}
	for key, tVal := range tMap {
		if sVal, ok := sMap[key]; !ok || tVal != sVal {
			return false
		}
	}
	return len(sMap) == len(tMap)
}
