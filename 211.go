package main

type WordDictionary struct {
	Chars     map[rune]*WordDictionary
	EndOfWord bool
}


func Constructor() WordDictionary {
	return WordDictionary{Chars: make(map[rune]*WordDictionary)}
}

func (this *WordDictionary) AddWord(word string) {
	for _, char := range word {
		dict, ok := this.Chars[char]
		if !ok {
			newDict := Constructor()
			this.Chars[char] = &newDict
			this = &newDict
			continue
		}
		this = dict
	}
	this.EndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	wordArr := []rune(word)
	var dfs func(wordSubArr []rune, root *WordDictionary) bool

	dfs = func(wordSubArr []rune, root *WordDictionary) bool {
		for i, char := range wordSubArr {
			if char == '.' {
				for _, branch := range root.Chars {
					if dfs(wordSubArr[i+1:], branch) {
						return true
					}
				}
				return false
			}
			dict, ok := root.Chars[char]
			if !ok {
				return false
			}
			root = dict
		}
		return root.EndOfWord
	}

	return dfs(wordArr, this)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
