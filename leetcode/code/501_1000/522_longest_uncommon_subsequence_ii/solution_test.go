package _22_longest_uncommon_subsequence_ii

import "testing"

func TestFindLUSLength(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{strs: []string{"aba", "cdc", "eae"}},
			want: 3,
		},
		{
			name: "second",
			args: args{strs: []string{"aa", "aaa", "aaa"}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLUSLength(tt.args.strs); got != tt.want {
				t.Errorf("FindLUSLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
