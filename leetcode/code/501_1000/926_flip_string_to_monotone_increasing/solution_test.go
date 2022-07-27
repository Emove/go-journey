package _26_flip_string_to_monotone_increasing

import "testing"

func TestMinFlipsMonoIncr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{s: "00110"},
			want: 1,
		},
		{
			name: "second",
			args: args{s: "010110"},
			want: 2,
		},
		{
			name: "third",
			args: args{s: "00011000"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinFlipsMonoIncr(tt.args.s); got != tt.want {
				t.Errorf("MinFlipsMonoIncr() = %v, want %v", got, tt.want)
			}
		})
	}
}
