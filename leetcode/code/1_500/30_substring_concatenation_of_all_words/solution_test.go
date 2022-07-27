package _0_substring_concatenation_of_all_words

import (
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{s: "barfoothefoobarman", words: []string{"foo", "bar"}},
			want: []int{0, 9},
		},
		{
			name: "second",
			args: args{s: "wordgoodgoodgoodbestwordd", words: []string{"word", "good", "best", "word"}},
			want: []int{},
		},
		{
			name: "third",
			args: args{s: "barfoofoobarthefoobarman", words: []string{"bar", "foo", "the"}},
			want: []int{6, 9, 12},
		},
		{
			name: "forth",
			args: args{s: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "good"}},
			want: []int{8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
