package main

import "fmt"

func main() {

	words := []string{"aa"}
	board := [][]byte{
		{'a', 'a'},
	}
	out := findWords(board, words)
	fmt.Printf("%v", out)
}
