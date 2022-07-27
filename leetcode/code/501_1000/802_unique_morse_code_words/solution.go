package _02_unique_morse_code_words

import (
	"strings"
)

var morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
	"....", "..", ".---", "-.-", ".-..", "--", "-.",
	"---", ".--.", "--.-", ".-.", "...", "-", "..-",
	"...-", ".--", "-..-", "-.--", "--.."}

func UniqueMorseRepresentations(words []string) int {
	set := make(map[string]struct{})

	for _, word := range words {
		builder := &strings.Builder{}
		for _, ch := range word {
			builder.WriteString(morse[ch-'a'])
		}
		set[builder.String()] = struct{}{}
	}

	return len(set)
}
