package _56_next_greater_element_iii

import "testing"

func TestNextGreaterElement(t *testing.T) {
	type args struct {
		args int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{args: 12},
			want: 21,
		},
		{
			name: "second",
			args: args{args: 21},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextGreaterElement(tt.args.args); got != tt.want {
				t.Errorf("NextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
