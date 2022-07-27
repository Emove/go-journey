package _68_validate_ip_address

import (
	"testing"
)

type args struct {
	queryIP string
}

type TE struct {
	name string
	args args
	want string
}

func TestValidIPAddress(t *testing.T) {
	tes := []TE{}
	tes = append(tes, getIPV4Testing()...)
	tes = append(tes, getIPV6Testing()...)
	for _, tt := range tes {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidIPAddress(tt.args.queryIP); got != tt.want {
				t.Errorf("ValidIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getIPV4Testing() []TE {
	return []TE{
		{
			name: "ipv4-first",
			args: args{queryIP: "172.16.254.1"},
			want: "IPv4",
		},
		{
			name: "ipv4-second",
			args: args{queryIP: "192.168.1.00"},
			want: "Neither",
		},
		{
			name: "ipv4-third",
			args: args{queryIP: "192.168@1.1"},
			want: "Neither",
		},
		{
			name: "ipv4-forth",
			args: args{queryIP: "192.168.1.0"},
			want: "IPv4",
		},
		{
			name: "ipv4-fifth",
			args: args{queryIP: "192.168.01.1"},
			want: "Neither",
		},
		{
			name: "ipv4-sixth",
			args: args{queryIP: "255.255.254.256"},
			want: "Neither",
		},
		{
			name: "ipv4-seventh",
			args: args{queryIP: "..133.34"},
			want: "Neither",
		},
	}
}

func getIPV6Testing() []TE {
	return []TE{
		{
			name: "ipv6-first",
			args: args{queryIP: "2001:0db8:85a3:0:0:8A2E:0370:7334"},
			want: "IPv6",
		},
		{
			name: "ipv6-second",
			args: args{queryIP: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"},
			want: "IPv6",
		},
		{
			name: "ipv6-third",
			args: args{queryIP: "2001:db8:85a3:0:0:8A2E:0370:7334"},
			want: "IPv6",
		},
		{
			name: "ipv6-forth",
			args: args{queryIP: "2001:0db8:85a3::8A2E:037j:7334"},
			want: "Neither",
		},
		{
			name: "ipv6-fifth",
			args: args{queryIP: "02001:0db8:85a3:0000:0000:8a2e:0370:7334"},
			want: "Neither",
		},
		{
			name: "ipv6-sixth",
			args: args{queryIP: "0::81:0:0:8:0:7"},
			want: "Neither",
		},
	}
}
