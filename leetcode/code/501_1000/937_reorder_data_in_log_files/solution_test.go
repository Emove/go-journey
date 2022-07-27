package _37_reorder_data_in_log_files

import (
	"reflect"
	"testing"
)

func TestReorderLogFiles(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "first",
			args: args{logs: []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}},
			want: []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
		},
		{
			name: "second",
			args: args{logs: []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}},
			want: []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReorderLogFiles1(tt.args.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReorderLogFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
