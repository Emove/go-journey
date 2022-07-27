package _24_goat_latin

import (
	"strings"
)

const VOWELS = "aeiouAEIOU"

func ToGoatLatin(sentence string) string {
	words := strings.Split(sentence, " ")
	res := strings.Builder{}
	suffix := strings.Builder{}
	for _, word := range words {
		suffix.WriteString("a")
		if strings.ContainsAny(VOWELS, strings.ToLower(string(word[0]))) {
			res.WriteString(word)
		} else {
			res.WriteString(word[1:])
			res.WriteByte(word[0])
		}
		res.WriteString("ma")
		res.WriteString(suffix.String())
		res.WriteString(" ")
	}
	s := res.String()
	return s[:len(s)-1]
}
