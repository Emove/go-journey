package combine_77

import (
	"fmt"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "first",
			args: args{n: 4, k: 2},
			want: [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}},
		},
		{
			name: "second",
			args: args{n: 1, k: 1},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combine(tt.args.n, tt.args.k)
			gotMap := make(map[string]byte, len(got))
			for i := 0; i < len(got); i++ {
				gotMap[serialize(got[i])] = 1
			}
			for i := 0; i < len(tt.want); i++ {
				if _, ok := gotMap[serialize(tt.want[i])]; !ok {
					t.Fatalf("combine() = %v, want %v", got, tt.want)
				}
			}
			t.Logf("combine() = %v, want %v", got, tt.want)
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
