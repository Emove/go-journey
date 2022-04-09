package _80_reaching_points

import "testing"

func TestReachingPoints(t *testing.T) {
	type args struct {
		sx int
		sy int
		tx int
		ty int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{sx: 1, sy: 1, tx: 3, ty: 5},
			want: true,
		},
		{
			name: "second",
			args: args{sx: 1, sy: 1, tx: 2, ty: 2},
			want: false,
		},
		{
			name: "third",
			args: args{sx: 1, sy: 1, tx: 1, ty: 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReachingPoints(tt.args.sx, tt.args.sy, tt.args.tx, tt.args.ty); got != tt.want {
				t.Errorf("ReachingPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
