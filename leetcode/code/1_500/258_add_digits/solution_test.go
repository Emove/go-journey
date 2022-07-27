package _58_add_digits

import "testing"

func TestAddDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{num: 38},
			want: 2,
		},
		{
			name: "second",
			args: args{num: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddDigits1(tt.args.num); got != tt.want {
				t.Errorf("AddDigits1() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddDigits2(tt.args.num); got != tt.want {
				t.Errorf("AddDigits1() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddDigits3(tt.args.num); got != tt.want {
				t.Errorf("AddDigits1() = %v, want %v", got, tt.want)
			}
		})
	}
}
