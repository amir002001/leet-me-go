package main

import (
	"strconv"
	"strings"
)

type Codec struct{}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	var builder strings.Builder
	for i := 0; i < len(strs); i++ {
		builder.WriteString(strconv.Itoa(len(strs[i])))
		builder.WriteRune('@')
		builder.WriteString(strs[i])
	}
	return builder.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	out := []string{}
	idx := 0
	var builder strings.Builder
	for idx < len(strs) {
		if _, err := strconv.Atoi(string(strs[idx])); err == nil {
			builder.WriteByte(strs[idx])
			idx++
		} else {
			idx++
			// go past @ sign
			length, err := strconv.Atoi(builder.String())
			if err != nil {
				panic(1)
			}
			out = append(out, strs[idx:idx+length])
			builder.Reset()
			idx += length
		}
	}
	return out
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
