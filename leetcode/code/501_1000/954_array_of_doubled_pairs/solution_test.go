package _54_array_of_doubled_pairs

import "testing"

func TestCanRecorderDoubled(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{arr: []int{1, 2, 4, 16, 8, 4}},
			want: false,
		},
		{
			name: "second",
			args: args{arr: []int{1, 2, 1, -8, 8, -4, 4, -4, 2, -2}},
			want: true,
		},
		{
			name: "third",
			args: args{arr: []int{4, -2, 2, -4}},
			want: true,
		},
		{
			name: "forth",
			args: args{arr: []int{3, 1, 3, 6}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanRecorderDoubled(tt.args.arr); got != tt.want {
				t.Errorf("CanRecorderDoubled() = %v, want %v", got, tt.want)
			}
		})
	}
}
