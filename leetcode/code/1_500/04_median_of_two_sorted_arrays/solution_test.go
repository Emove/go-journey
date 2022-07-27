package _4_median_of_two_sorted_arrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "first",
			args: args{nums1: []int{1, 3}, nums2: []int{2}},
			want: float64(2),
		},
		{
			name: "second",
			args: args{nums1: []int{1, 2}, nums2: []int{3, 4}},
			want: 2.5,
		},
		{
			name: "third",
			args: args{nums1: []int{}, nums2: []int{1}},
			want: float64(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("FindMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
