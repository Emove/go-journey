package _217_minimum_cost_to_move_chips_to_the_same_position

import "testing"

func TestMinCostToMoveChips(t *testing.T) {
	type args struct {
		position []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{position: []int{1, 2, 3}},
			want: 1,
		},
		{
			name: "second",
			args: args{position: []int{2, 2, 2, 3, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCostToMoveChips(tt.args.position); got != tt.want {
				t.Errorf("MinCostToMoveChips() = %v, want %v", got, tt.want)
			}
		})
	}
}
