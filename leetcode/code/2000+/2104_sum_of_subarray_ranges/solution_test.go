package sum_of_subarray_ranges_2104

import "testing"

func Test_subArrayRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name:    "first",
			args:    args{nums: []int{1, 2, 3}},
			wantAns: 4,
		},
		{
			name:    "second",
			args:    args{nums: []int{1, 3, 3}},
			wantAns: 4,
		},
		{
			name:    "third",
			args:    args{nums: []int{4, -2, -3, 4, 1}},
			wantAns: 59,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := SubArrayRanges(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("subArrayRanges() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := SubArrayRanges1(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("subArrayRanges() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
