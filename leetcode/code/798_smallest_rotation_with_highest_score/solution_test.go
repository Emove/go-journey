package _98_smallest_rotation_with_highest_score

import "testing"

func TestBestRotation(t *testing.T) {
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
			args: args{nums: []int{2, 3, 1, 4, 0}},
			want: 3,
		},
		{
			name: "second",
			args: args{nums: []int{1, 3, 0, 2, 4}},
			want: 0,
		},
		{
			name: "third",
			args: args{nums: []int{6, 2, 8, 3, 5, 2, 4, 3, 7, 6}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BestRotation(tt.args.nums); got != tt.want {
				t.Errorf("BestRotation() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BestRotation1(tt.args.nums); got != tt.want {
				t.Errorf("BestRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
