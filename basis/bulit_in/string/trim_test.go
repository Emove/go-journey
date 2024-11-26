package string

import "testing"

func TestTrimSuffix(t *testing.T) {
	type args struct {
		s      string
		suffix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{
				s:      "hello_world.log",
				suffix: ".log",
			},
			want: "hello_world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimSuffix(tt.args.s, tt.args.suffix); got != tt.want {
				t.Errorf("TrimSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
