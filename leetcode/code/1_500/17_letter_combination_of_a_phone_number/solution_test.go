package _7_letter_combination_of_a_phone_number

import (
	"leetcode/code/eq"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "first",
			args: args{digits: "23"},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "second",
			args: args{digits: ""},
			want: []string{},
		},
		{
			name: "third",
			args: args{digits: "2"},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LetterCombinations2(tt.args.digits); !eq.EqualSliceString(got, tt.want) {
				t.Errorf("LetterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
