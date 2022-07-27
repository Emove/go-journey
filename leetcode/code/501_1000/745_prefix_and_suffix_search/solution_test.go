package _45_prefix_and_suffix_search

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	wordFilter := Constructor([]string{"abbba", "abba"})
	fmt.Println(wordFilter.F("ab", "ba"))
	//fmt.Println(wordFilter.F("a", "a"))
}
