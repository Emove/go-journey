package _42_di_string_match

import (
	"leetcode/code/eq"
	"testing"
)

func TestDiStringMatch(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{s: "IDID"},
			want: []int{0, 4, 1, 3, 2},
		},
		{
			name: "second",
			args: args{s: "III"},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "third",
			args: args{s: "DDI"},
			want: []int{3, 2, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiStringMatch(tt.args.s); !eq.EqualSliceInt(got, tt.want) {
				t.Errorf("DiStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
