package _88_longest_absolute_file_path

import "testing"

func TestLengthLongestPath(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{input: "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"},
			want: 20,
		},
		{
			name: "second",
			args: args{input: "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"},
			want: 32,
		},
		{
			name: "third",
			args: args{input: "a"},
			want: 0,
		},
		{
			name: "forth",
			args: args{input: "file1.txt\nfile2.txt\nlongfile.txt"},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthLongestPath(tt.args.input); got != tt.want {
				t.Errorf("LengthLongestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
