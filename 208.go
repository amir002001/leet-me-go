package main

type Trie struct {
	Chars     map[rune]*Trie
	EndOfWord bool
}

func DudeChamp() Trie {
	return Trie{make(map[rune]*Trie), false}
}
func (this *Trie) Insert(word string) {
	for _, char := range word {
		trie, ok := this.Chars[char]
		if !ok {
			newTrie := DudeChamp()
			this.Chars[char] = &newTrie
			this = &newTrie
			continue
		}
		this = trie
	}
	this.EndOfWord = true
}

func (this *Trie) Search(word string) bool {
	for _, char := range word {
		trie, ok := this.Chars[char]
		if !ok {
			return false
		}
		this = trie
	}
	return this.EndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, char := range prefix {
		trie, ok := this.Chars[char]
		if !ok {
			return false
		}
		this = trie
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
