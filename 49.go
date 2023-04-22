package main

func groupAnagrams(strs []string) [][]string {
	freqMap := make(map[[26]int][]string)
	for _, str := range strs {
		freq := getStringFreq(str)
		freqMap[freq] = append(freqMap[freq], str)
	}
	output := [][]string{}
	for _, arr := range freqMap {
		output = append(output, arr)
	}
	return output
}

func getStringFreq(s string) [26]int {
	out := [26]int{}
	for _, char := range s {
		out[char-'a']++
	}

	return out
}
