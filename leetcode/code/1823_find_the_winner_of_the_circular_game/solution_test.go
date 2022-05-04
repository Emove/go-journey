package _823_find_the_winner_of_the_circular_game

import "testing"

func TestFindTheWinner(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name:    "first",
			args:    args{n: 5, k: 2},
			wantRes: 3,
		},
		{
			name:    "second",
			args:    args{n: 6, k: 5},
			wantRes: 1,
		},
		{
			name:    "third",
			args:    args{n: 3, k: 1},
			wantRes: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := FindTheWinner1(tt.args.n, tt.args.k); gotRes != tt.wantRes {
				t.Errorf("FindTheWinner() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
