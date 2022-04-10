package _02_unique_morse_code_words

import "testing"

func TestUniqueMorseRepresentations(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{words: []string{"gin", "zen", "gig", "msg"}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueMorseRepresentations(tt.args.words); got != tt.want {
				t.Errorf("UniqueMorseRepresentations() = %v, want %v", got, tt.want)
			}
		})
	}
}
