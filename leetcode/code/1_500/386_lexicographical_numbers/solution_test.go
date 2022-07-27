package _86_lexicographical_numbers

import (
	"reflect"
	"testing"
)

func TestLexicalOrder(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{n: 13},
			want: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "second",
			args: args{n: 2},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LexicalOrder(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LexicalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
