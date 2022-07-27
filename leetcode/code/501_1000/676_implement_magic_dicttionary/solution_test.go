package _76_implement_magic_dicttionary

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	dictionary := Constructor()
	dictionary.BuildDict([]string{"hello", "leetcode"})
	//t.Logf("search: hello, want: %v, got: %v", false, dictionary.Search("hello"))
	t.Logf("search: hhllo, want: %v, got: %v", true, dictionary.Search("hhllo"))
	//t.Logf("search: hell, want: %v, got: %v", false, dictionary.Search("hell"))
	//t.Logf("search: leetcoded, want: %v, got: %v", false, dictionary.Search("leetcoded"))
}
