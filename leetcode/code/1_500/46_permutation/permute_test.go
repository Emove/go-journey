package permutation_46

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "first",
			args: args{nums: []int{1, 2, 3}},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "second",
			args: args{nums: []int{0, 1}},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "third",
			args: args{nums: []int{1}},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.args.nums)
			want := tt.want
			gotMap := make(map[string]byte, len(got)*len(got[0]))
			for i := 0; i < len(got); i++ {
				gotMap[serialize(got[i])] = 1
			}
			for _, nums := range want {
				if _, ok := gotMap[serialize(nums)]; !ok {
					t.Fatalf("permute() = %v, want %v", got, tt.want)
				}
			}
			t.Logf("permute() = %v, want %v", got, tt.want)
		})
	}
}

func serialize(nums []int) string {
	res := ""
	n := len(nums)
	for i := 0; i < n; i++ {
		res += fmt.Sprintf("%d", nums[i])
		if i < n-1 {
			res += "_"
		}
	}
	return res
}
