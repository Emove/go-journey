package _19_most_common_word

import "testing"

func TestMostCommonWord(t *testing.T) {
	type args struct {
		paragraph string
		banned    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.", banned: []string{"hit"}},
			want: "ball",
		},
		{
			name: "second",
			args: args{paragraph: "Bob", banned: []string{}},
			want: "bob",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MostCommonWord(tt.args.paragraph, tt.args.banned); got != tt.want {
				t.Errorf("MostCommonWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
