package _108_defanging_an_ip_address

import "testing"

func TestDefangIPaddr(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{address: "1.1.1.1"},
			want: "1[.]1[.]1[.]1",
		},
		{
			name: "second",
			args: args{address: "255.100.50.0"},
			want: "255[.]100[.]50[.]0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefangIPaddr(tt.args.address); got != tt.want {
				t.Errorf("DefangIPaddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
