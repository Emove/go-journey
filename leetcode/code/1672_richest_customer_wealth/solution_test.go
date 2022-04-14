package _672_richest_customer_wealth

import "testing"

func TestMaximumWealth(t *testing.T) {
	type args struct {
		accounts [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{accounts: [][]int{{1, 2, 3}, {3, 2, 1}}},
			want: 6,
		},
		{
			name: "second",
			args: args{accounts: [][]int{{1, 5}, {3, 5}, {7, 3}}},
			want: 10,
		},
		{
			name: "third",
			args: args{accounts: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumWealth(tt.args.accounts); got != tt.want {
				t.Errorf("MaximumWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
