package file

import "testing"

func TestWalk(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "first",
			args: args{
				path: "./configs",
			},
			wantErr: false,
		},
		{
			name: "second",
			args: args{
				path: "./",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Walk(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Walk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParent(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{
				path: "D:\\workspace\\profile\\go\\go-journey\\file\\walk.go",
			},
			want: "D:\\workspace\\profile\\go\\go-journey",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parent(tt.args.path); got != tt.want {
				t.Errorf("Parent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCurrentDir(t *testing.T) {
	dir, err := CurrentDir()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("dir: %s", dir)
}
