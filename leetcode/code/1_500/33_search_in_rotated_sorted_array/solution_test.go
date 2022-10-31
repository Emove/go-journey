package _3_search_in_rotated_sorted_array

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0},
			want: 4,
		},
		{
			name: "second",
			args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3},
			want: -1,
		},
		{
			name: "third",
			args: args{nums: []int{1}, target: 0},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
