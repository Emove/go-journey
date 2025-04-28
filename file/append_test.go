package file

import (
	"testing"
)

func TestOpenWithNotExist(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestOpenWithNotExist",
			args: args{
				name: "testfile",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := OpenWithNotExist(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenWithNotExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAppend(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := Append("testfile", []byte("test data"))
		if err != nil {
			return
		}
	}
}
