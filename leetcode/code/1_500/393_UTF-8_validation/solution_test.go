package _93_UTF_8_validation

import (
	"fmt"
	"testing"
)

func TestValidUtf8(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{data: []int{197, 130, 1}},
			want: true,
		},
		{
			name: "second",
			args: args{data: []int{235, 140, 4}},
			want: false,
		},
		{
			name: "third",
			args: args{data: []int{240, 162, 138, 147, 145}},
			want: false,
		},
		{
			name: "forth",
			args: args{data: []int{250, 145, 145, 145, 145}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidUtf8(tt.args.data); got != tt.want {
				t.Errorf("ValidUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	fmt.Println(1 & 192)
}
