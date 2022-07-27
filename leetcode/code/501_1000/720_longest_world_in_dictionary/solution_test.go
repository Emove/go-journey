package _20_longest_world_in_dictionary

import "testing"

func TestLongestWorld(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{words: []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}},
			want: "apple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestWorld(tt.args.words); got != tt.want {
				t.Errorf("LongestWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}
