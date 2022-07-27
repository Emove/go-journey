package _1_05_one_away_lcci

import (
	"testing"
)

func TestOneEditAway(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{first: "pale", second: "ple"},
			want: true,
		},
		{
			name: "second",
			args: args{first: "pales", second: "pal"},
			want: false,
		},
		{
			name: "third",
			args: args{first: "", second: "b"},
			want: true,
		},
		{
			name: "forth",
			args: args{first: "mart", second: "karma"},
			want: false,
		},
		{
			name: "fifth",
			args: args{first: "a", second: "b"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OneEditAway(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("OneEditAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
