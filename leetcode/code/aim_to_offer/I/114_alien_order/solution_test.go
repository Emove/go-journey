package _14_alien_order

import "testing"

func TestAlienOrder(t *testing.T) {
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
			args: args{words: []string{"wrt", "wrf", "er", "ett", "rftt"}},
			want: "wertf",
		},
		{
			name: "second",
			args: args{[]string{"z", "x"}},
			want: "zx",
		},
		{
			name: "third",
			args: args{words: []string{"z", "x", "z"}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AlienOrder(tt.args.words); got != tt.want {
				t.Errorf("AlienOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
