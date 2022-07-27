package _028_find_missing_observations

import (
	"reflect"
	"testing"
)

func TestMissingRolls(t *testing.T) {
	type args struct {
		rolls []int
		mean  int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{rolls: []int{3, 2, 4, 3}, mean: 4, n: 2},
			want: []int{6, 6},
		},
		{
			name: "second",
			args: args{rolls: []int{1, 5, 6}, mean: 3, n: 4},
			want: []int{3, 2, 2, 2},
		},
		{
			name: "third",
			args: args{rolls: []int{1, 2, 3, 4}, mean: 6, n: 4},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MissingRolls(tt.args.rolls, tt.args.mean, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MissingRolls() = %v, want %v", got, tt.want)
			}
		})
	}
}
