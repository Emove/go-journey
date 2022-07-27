package decode_xored_array_1720

import (
	"reflect"
	"testing"
)

func Test_decode(t *testing.T) {
	type args struct {
		encode []int
		first  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				encode: []int{1, 2, 3},
				first:  1,
			},
			want: []int{1, 0, 2, 1},
		},
		{
			name: "second",
			args: args{
				encode: []int{6, 2, 7, 3},
				first:  4,
			},
			want: []int{4, 2, 0, 7, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.encode, tt.args.first); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
