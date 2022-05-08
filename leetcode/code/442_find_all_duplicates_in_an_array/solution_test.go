package _42_find_all_duplicates_in_an_array

import (
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}},
			want: []int{2, 3},
		},
		{
			name: "second",
			args: args{nums: []int{1, 1, 2}},
			want: []int{1},
		},
		{
			name: "third",
			args: args{nums: []int{1}},
			want: []int{},
		},
		{
			name: "forth",
			args: args{nums: []int{5, 4, 6, 7, 9, 3, 10, 9, 5, 6}},
			want: []int{9, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicates(tt.args.nums); !eq(got, tt.want) {
				t.Errorf("FindDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := FindDuplicates1(tt.args.nums); !eq(got, tt.want) {
	//			t.Errorf("FindDuplicates() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

func eq(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	nums1Map := make(map[int]struct{}, len(nums1))
	for _, v := range nums1 {
		nums1Map[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := nums1Map[v]; !ok {
			return false
		}
	}
	return true
}
