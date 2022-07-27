package _29_unique_email_addresses

import "testing"

func TestNumUniqueEmails(t *testing.T) {
	type args struct {
		emails []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{emails: []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}},
			want: 2,
		},
		{
			name: "second",
			args: args{emails: []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumUniqueEmails(tt.args.emails); got != tt.want {
				t.Errorf("NumUniqueEmails() = %v, want %v", got, tt.want)
			}
		})
	}
}
