package file

import (
	"testing"
)

func TestState(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		args    args
		wantNil bool
		wantErr bool
	}{
		{
			args:    args{path: "./config.yaml"},
			wantNil: false,
			wantErr: false,
		},
		{
			args:    args{path: "./configs/config.yaml"},
			wantNil: false,
			wantErr: false,
		},
		{
			args:    args{path: "D:\\workspace\\profile\\go\\go-journey\\file\\config.yaml"},
			wantNil: false,
			wantErr: false,
		},
		{
			args:    args{path: "D:\\workspace\\profile\\go\\go-journey\\file\\configs\\config.yaml"},
			wantNil: false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := State(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("State() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantNil && got == nil {
				t.Errorf("State() FileInfo = %v, wantNil %v", got, tt.wantNil)
				return
			}
		})
	}
}
