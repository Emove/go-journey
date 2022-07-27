package _90_find_and_replace_pattern

import (
	"leetcode/code/eq"
	"testing"
)

func TestFindAndReplacePattern(t *testing.T) {
	type args struct {
		words   []string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "first",
			args: args{words: []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, pattern: "abb"},
			want: []string{"mee", "aqq"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAndReplacePattern(tt.args.words, tt.args.pattern); !eq.EqualSliceString(got, tt.want) {
				t.Errorf("FindAndReplacePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
