package _04_base_7

import "testing"

func TestConvertToBase7(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{num: 100},
			want: "202",
		},
		{
			name: "second",
			args: args{num: -7},
			want: "-10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToBase7(tt.args.num); got != tt.want {
				t.Errorf("ConvertToBase7() = %v, want %v", got, tt.want)
			}
		})
	}
}
