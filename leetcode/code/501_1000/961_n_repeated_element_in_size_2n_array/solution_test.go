package _61_n_repeated_element_in_size_2n_array

import "testing"

func TestRepeatedNTimes(t *testing.T) {
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
			args: args{nums: []int{1, 2, 3, 3}},
			want: 3,
		},
		{
			name: "second",
			args: args{nums: []int{2, 1, 2, 5, 3, 2}},
			want: 2,
		},
		{
			name: "third",
			args: args{nums: []int{5, 1, 5, 2, 5, 3, 5, 4}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepeatedNTimes(tt.args.nums); got != tt.want {
				t.Errorf("RepeatedNTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
