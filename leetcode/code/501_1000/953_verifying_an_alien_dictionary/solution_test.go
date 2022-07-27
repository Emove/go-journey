package _53_verifying_an_alien_dictionary

import "testing"

func TestIsAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{words: []string{"hello", "leetcode"}, order: "hlabcdefgijkmnopqrstuvwxyz"},
			want: true,
		},
		{
			name: "second",
			args: args{words: []string{"word", "world", "row"}, order: "worldabcefghijkmnpqstuvxyz"},
			want: false,
		},
		{
			name: "third",
			args: args{words: []string{"apple", "app"}, order: "abcdefghijklmnopqrstuvwxyz"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlienSorted(tt.args.words, tt.args.order); got != tt.want {
				t.Errorf("IsAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
