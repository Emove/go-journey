package replace_all_question_mark_1576

import (
	"testing"
)

func Test_modifyString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{s: "?zs"},
			want: "azs",
		},
		{
			name: "second",
			args: args{s: "ubv?w"},
			want: "ubvaw",
		},
		{
			name: "third",
			args: args{s: "ubv?w"},
			want: "ubvaw",
		},
		{
			name: "forth",
			args: args{s: "j?qg??b"},
			want: "jaqgacb",
		},
		{
			name: "fifth",
			args: args{s: "??yw?ipkj?"},
			want: "abywaipkja",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifyString(tt.args.s); got != tt.want {
				t.Errorf("modifyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
