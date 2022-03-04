package sorted_squares_977

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{[]int{-4, -1, 0, 3, 10}},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "second",
			args: args{[]int{-7, -3, 2, 3, 11}},
			want: []int{4, 9, 9, 49, 121},
		},
		{
			name: "third",
			args: args{[]int{0, 2}},
			want: []int{0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
