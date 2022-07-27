package length_of_longest_substrings_3

import "testing"

func Test_lengthOfLongestSubString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "second",
			args: args{s: "bbbbb"},
			want: 1,
		},
		{
			name: "third",
			args: args{s: "pwwkew"},
			want: 3,
		},
		{
			name: "forth",
			args: args{s: ""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubString(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubString() = %v, want %v", got, tt.want)
			}
		})
	}
}
