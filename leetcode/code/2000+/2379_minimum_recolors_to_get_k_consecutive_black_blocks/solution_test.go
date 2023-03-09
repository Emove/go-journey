package _379_minimum_recolors_to_get_k_consecutive_black_blocks

import "testing"

func Test_minimumRecolors(t *testing.T) {
	type args struct {
		blocks string
		k      int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{blocks: "WBBWWBBWBW", k: 7},
			want: 3,
		},
		{
			args: args{blocks: "WBWBBBW", k: 2},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minimumRecolors(tt.args.blocks, tt.args.k); got != tt.want {
				t.Errorf("minimumRecolors() = %v, want %v", got, tt.want)
			}
		})
	}
}
