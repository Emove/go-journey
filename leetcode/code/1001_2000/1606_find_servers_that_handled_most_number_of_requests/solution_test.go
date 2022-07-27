package _606_find_servers_that_handled_most_number_of_requests

import (
	"reflect"
	"testing"
)

func Test_busiestServers(t *testing.T) {
	type args struct {
		k       int
		arrival []int
		load    []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name:    "first",
			args:    args{k: 3, arrival: []int{1, 2, 3, 4, 5}, load: []int{5, 2, 3, 3, 3}},
			wantAns: []int{1},
		},
		{
			name:    "second",
			args:    args{k: 3, arrival: []int{1, 2, 3, 4}, load: []int{1, 2, 1, 2}},
			wantAns: []int{0},
		},
		{
			name:    "third",
			args:    args{k: 3, arrival: []int{1, 2, 3}, load: []int{10, 12, 11}},
			wantAns: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := busiestServers(tt.args.k, tt.args.arrival, tt.args.load); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("busiestServers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
