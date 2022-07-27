package _61_image_smoother

import (
	"reflect"
	"testing"
)

func TestImageSmoother(t *testing.T) {
	type args struct {
		img [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "first",
			args: args{img: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
			want: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			name: "second",
			args: args{img: [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}},
			want: [][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ImageSmoother2(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImageSmoother() = %v, want %v", got, tt.want)
			}
		})
	}
}
