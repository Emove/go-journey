package _044_count_number_of_maximum_bitwise_or_subsets

import "testing"

func Test_countMaxOrSubsets(t *testing.T) {
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
			args: args{nums: []int{3, 1}},
			want: 2,
		},
		{
			name: "second",
			args: args{nums: []int{2, 2, 2}},
			want: 7,
		},
		{
			name: "third",
			args: args{nums: []int{3, 2, 1, 5}},
			want: 6,
		},
	}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := countMaxOrSubsets(tt.args.nums); got != tt.want {
	//			t.Errorf("countMaxOrSubsets() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}

	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := dp(tt.args.nums); got != tt.want {
	//			t.Errorf("countMaxOrSubsets() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dfs(tt.args.nums); got != tt.want {
				t.Errorf("countMaxOrSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
