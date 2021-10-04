package letter_case_permutation_784

import (
	"testing"
)

func Test_letterCasePermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "first",
			args: args{s: "a1b2"},
			want: []string{"a1b2", "a1B2", "A1b2", "A1B2"},
		},
		{
			name: "second",
			args: args{s: "3z4"},
			want: []string{"3z4", "3Z4"},
		},
		{
			name: "third",
			args: args{s: "12345"},
			want: []string{"12345"},
		},
		{
			name: "forth",
			args: args{s: "0"},
			want: []string{"0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCasePermutation(tt.args.s)
			want := tt.want
			gotMap := make(map[string]byte, len(got))
			for _, val := range got {
				gotMap[val] = 1
			}
			for _, s := range want {
				if _, ok := gotMap[s]; !ok {
					t.Fatalf("letterCasePermutation() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
