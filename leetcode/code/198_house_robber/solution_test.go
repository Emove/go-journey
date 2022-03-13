package _98_house_robber

import "testing"

func TestRob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{nums: []int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "second",
			args: args{nums: []int{2, 7, 9, 3, 1}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rob(tt.args.nums); got != tt.want {
				t.Errorf("Rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
