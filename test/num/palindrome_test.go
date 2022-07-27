package num

import "testing"

func TestVerifyPalindrome(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{n: 99},
			want: true,
		},
		{
			name: "second",
			args: args{n: 1},
			want: true,
		},
		{
			name: "third",
			args: args{n: 24},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.n); got != tt.want {
				t.Errorf("VerifyPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
