package _67_unique_substrings_in_wraparound_string

import "testing"

func TestFindSubstringInWraproundString(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{p: "a"},
			want: 1,
		},
		{
			name: "second",
			args: args{p: "cac"},
			want: 2,
		},
		{
			name: "third",
			args: args{p: "zab"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSubstringInWraproundString(tt.args.p); got != tt.want {
				t.Errorf("FindSubstringInWraproundString() = %v, want %v", got, tt.want)
			}
		})
	}
}
