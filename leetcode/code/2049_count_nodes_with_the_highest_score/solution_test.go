package _049_count_nodes_with_the_highest_score

import "testing"

func TestCountHighestScoreNodes(t *testing.T) {
	type args struct {
		parents []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{parents: []int{-1, 2, 0, 2, 0}},
			want: 3,
		},
		{
			name: "second",
			args: args{parents: []int{-1, 2, 0}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountHighestScoreNodes(tt.args.parents); got != tt.want {
				t.Errorf("CountHighestScoreNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
