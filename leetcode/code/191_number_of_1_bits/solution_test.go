package _91_number_of_1_bits

import "testing"

func TestHammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name     string
		args     args
		wantOnes int
	}{
		{
			name:     "first",
			args:     args{num: 5},
			wantOnes: 2,
		},
		{
			name:     "second",
			args:     args{num: 6},
			wantOnes: 2,
		},
		{
			name:     "third",
			args:     args{num: 7},
			wantOnes: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOnes := HammingWeight(tt.args.num); gotOnes != tt.wantOnes {
				t.Errorf("HammingWeight() = %v, want %v", gotOnes, tt.wantOnes)
			}
		})
	}
}
