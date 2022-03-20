package _039_the_time_when_the_network_becomes_idle

import "testing"

func TestNetworkBecomesIdle(t *testing.T) {
	type args struct {
		edges    [][]int
		patience []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "first",
			args:    args{edges: [][]int{{0, 1}, {1, 2}}, patience: []int{0, 2, 1}},
			wantAns: 8,
		},
		{
			name:    "first",
			args:    args{edges: [][]int{{0, 1}, {1, 2}, {0, 2}}, patience: []int{0, 10, 10}},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := NetworkBecomesIdle(tt.args.edges, tt.args.patience); gotAns != tt.wantAns {
				t.Errorf("NetworkBecomesIdle() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
