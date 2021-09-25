package reverse_strings_344

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{s: []byte{'h', 'e', 'l', 'l', 'o'}},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name: "second",
			args: args{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if reverseString(tt.args.s); !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("reverseString() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
