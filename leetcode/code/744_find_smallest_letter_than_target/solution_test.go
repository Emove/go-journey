package _44_find_smallest_letter_than_target

import "testing"

func TestNextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "first",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'a'},
			want: 'c',
		},
		{
			name: "second",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'c'},
			want: 'f',
		}, {
			name: "third",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'd'},
			want: 'f',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("NextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
