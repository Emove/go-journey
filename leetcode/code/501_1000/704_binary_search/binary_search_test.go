package _04_binary_search

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intervalSearch(t *testing.T) {
	type args struct {
		nums      []int
		target    int
		isGreater bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				nums:      []int{-1, 0, 3, 5, 9, 12},
				target:    6,
				isGreater: true,
			},
			want: 4,
		},
		{
			name: "second",
			args: args{
				nums:      []int{-1, 0, 3, 5, 9, 12},
				target:    6,
				isGreater: false,
			},
			want: 3,
		},
		{
			name: "third",
			args: args{
				nums:      []int{-1, 0, 3, 5, 9, 12},
				target:    5,
				isGreater: false,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intervalSearch(tt.args.nums, tt.args.target, tt.args.isGreater); got != tt.want {
				t.Errorf("intervalSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
