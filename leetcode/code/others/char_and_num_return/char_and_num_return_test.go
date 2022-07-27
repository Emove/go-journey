package char_and_num_return

import "testing"

func Test_char_and_num_return(t *testing.T) {
	type args struct {
		text_source string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{text_source: "xb 1 cc 5 dd 10 ee 2"},
			want: "xb cc dd ee 1 2 5 10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := char_and_num_return(tt.args.text_source); got != tt.want {
				t.Errorf("char_and_num_return() = %v, want %v", got, tt.want)
			}
		})
	}
}
