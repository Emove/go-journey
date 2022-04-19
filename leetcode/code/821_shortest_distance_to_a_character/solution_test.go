package _21_shortest_distance_to_a_character

import (
	"reflect"
	"testing"
)

func TestShortestToChar(t *testing.T) {
	type args struct {
		s string
		c byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{s: "loveleetcode", c: 'e'},
			want: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
		{
			name: "second",
			args: args{s: "aaab", c: 'b'},
			want: []int{3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortestToChar(tt.args.s, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShortestToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
