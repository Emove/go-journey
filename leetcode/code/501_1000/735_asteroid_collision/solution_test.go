package _35_asteroid_collision

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{asteroids: []int{5, 10, -5}},
			want: []int{5, 10},
		},
		{
			name: "second",
			args: args{asteroids: []int{8, -8}},
			want: []int{},
		},
		{
			name: "third",
			args: args{asteroids: []int{-2, -1, 1, 2}},
			want: []int{-2, -1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsteroidCollision(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
