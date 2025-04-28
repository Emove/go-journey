package file

import "testing"

func TestGetDir(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{name: "./configs/config.yaml"},
			want: "./configs",
		},
		{
			args: args{name: "D:\\workspace\\profile\\go\\go-journey\\file\\config.yaml"},
			want: "D:\\workspace\\profile\\go\\go-journey\\file",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDir(tt.args.name); got != tt.want {
				t.Errorf("GetDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
