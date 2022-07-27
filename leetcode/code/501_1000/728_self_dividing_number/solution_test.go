package _28_self_dividing_number

import (
	"reflect"
	"testing"
)

func TestSelfDividingNumbers(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{left: 1, right: 22},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
		{
			name: "second",
			args: args{left: 47, right: 85},
			want: []int{48, 55, 66, 77},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelfDividingNumbers(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelfDividingNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
