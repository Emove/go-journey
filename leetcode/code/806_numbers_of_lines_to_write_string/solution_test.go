package _06_numbers_of_lines_to_write_string

import (
	"reflect"
	"testing"
)

func TestNumberOfLines(t *testing.T) {
	type args struct {
		widths []int
		s      string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{widths: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, s: "abcdefghijklmnopqrstuvwxyz"},
			want: []int{3, 60},
		},
		{
			name: "first",
			args: args{widths: []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, s: "bbbcccdddaaa"},
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfLines(tt.args.widths, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumberOfLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
