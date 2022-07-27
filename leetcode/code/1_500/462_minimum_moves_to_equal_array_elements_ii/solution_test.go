package _62_minimum_moves_to_equal_array_elements_ii

import "testing"

func TestMinMoves2(t *testing.T) {
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
			args: args{nums: []int{1, 2, 3}},
			want: 2,
		},
		{
			name: "second",
			args: args{nums: []int{1, 10, 2, 9}},
			want: 16,
		},
	}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := MinMoves2(tt.args.nums); got != tt.want {
	//			t.Errorf("MinMoves2() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMoves2i(tt.args.nums); got != tt.want {
				t.Errorf("MinMoves2() = %v, want %v", got, tt.want)
			}
		})
	}
}
