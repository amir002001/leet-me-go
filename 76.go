package main

func minWindow(s string, t string) string {
	found := false
	start, end := 0, len(s)-1
	requiredFreq := make(map[rune]int)
	currFreq := make(map[rune]int)
	freqMet := 0
	j := 0

	for _, char := range t {
		requiredFreq[char]++
	}

	for i, char := range s {
		if req, ok := requiredFreq[char]; ok {
			currFreq[char]++
			if currFreq[char] == req {
				freqMet++
			}
			for freqMet == len(requiredFreq) {
				found = true
				if (i - j + 1) < (end - start + 1) {
					end = i
					start = j
				}
				jRune := rune(s[j])
				if jReq, jOk := requiredFreq[jRune]; jOk {
					if currFreq[jRune] == jReq {
						freqMet--
					}
					currFreq[jRune]--
				}
				j++
			}
		}
	}

	if !found {
		return ""
	}
	return s[start : end+1]
}
