package _024_maximize_the_confusion_of_an_exam

import "testing"

func TestMaxConsecutiveAnswers(t *testing.T) {
	type args struct {
		answerKey string
		k         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{answerKey: "TTFF", k: 2},
			want: 4,
		},
		{
			name: "second",
			args: args{answerKey: "TFFT", k: 1},
			want: 3,
		},
		{
			name: "third",
			args: args{answerKey: "FFFTTFTTFT", k: 3},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxConsecutiveAnswers(tt.args.answerKey, tt.args.k); got != tt.want {
				t.Errorf("MaxConsecutiveAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}
