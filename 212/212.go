package main

import (
	"fmt"
)

type Trie struct {
	chars     map[rune]*Trie
	endOfWord bool
}

func findWords(board [][]byte, words []string) []string {
	trie := &Trie{make(map[rune]*Trie), false}
	for _, word := range words {
		insertWord(trie, word)
	}
	out := make(map[string]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if _, ok := trie.chars[rune(board[i][j])]; ok {
				visited := make(map[string]bool)
				dfsChar(board, trie, i, j, &out, &visited, "")
			}
		}
	}

	outSlice := make([]string, len(out))

	idx := 0
	for k := range out {
		outSlice[idx] = k
		idx++
	}
	return outSlice
}

// left here. we checked bounds now we need to add to visited, check the trie, if not check for eow, add it, delete from visited and done.
func dfsChar(
	board [][]byte,
	trie *Trie,
	i, j int,
	out *map[string]bool,
	visited *map[string]bool,
	word string,
) {
	if !boundsOk(i, j, board) {
		return
	}
	key := fmt.Sprintf("%d;%d", i, j)
	// if already visited leave
	if _, ok := (*visited)[key]; ok {
		return
	}
	(*visited)[key] = true
	// at the end, remove from visited
	defer delete((*visited), key)
	nextTrie, ok := trie.chars[rune(board[i][j])]
	if !ok {
		return
	}
	word += string(board[i][j])
	if nextTrie.endOfWord {
		(*out)[word] = true
	}

	dfsChar(board, nextTrie, i-1, j, out, visited, word)
	dfsChar(board, nextTrie, i+1, j, out, visited, word)
	dfsChar(board, nextTrie, i, j+1, out, visited, word)
	dfsChar(board, nextTrie, i, j-1, out, visited, word)
}

func boundsOk(i, j int, board [][]byte) bool {
	if i < 0 || i >= len(board) {
		return false
	}
	if j < 0 || j >= len(board[0]) {
		return false
	}
	return true
}

func insertWord(trie *Trie, word string) {
	for _, char := range word {
		if subTrie, ok := trie.chars[char]; ok {
			trie = subTrie
		} else {
			subTrie := &Trie{make(map[rune]*Trie), false}
			trie.chars[char] = subTrie
			trie = subTrie
		}
	}
	trie.endOfWord = true
}
