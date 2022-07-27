package _038_remove_colored_pieces_if_both_neighbors_are_the_same_color

import "testing"

func TestWinnerOfGame(t *testing.T) {
	type args struct {
		colors string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{colors: "AAABABB"},
			want: true,
		},
		{
			name: "second",
			args: args{colors: "AA"},
			want: false,
		},
		{
			name: "third",
			args: args{colors: "ABBBBBBBAAA"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WinnerOfGame(tt.args.colors); got != tt.want {
				t.Errorf("WinnerOfGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
