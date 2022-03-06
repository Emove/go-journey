package _001_good_day_to_rob_bank

import (
	"reflect"
	"testing"
)

func TestGoodDayToRobBank(t *testing.T) {
	type args struct {
		security []int
		time     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{security: []int{5, 3, 3, 3, 5, 6, 2}, time: 2},
			want: []int{2, 3},
		},
		{
			name: "second",
			args: args{security: []int{1, 1, 1, 1, 1}, time: 0},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "third",
			args: args{security: []int{1, 2, 3, 4, 5, 6}, time: 2},
			want: []int{},
		},
		{
			name: "forth",
			args: args{security: []int{1}, time: 5},
			want: []int{},
		},
	}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := GoodDayToRobBank(tt.args.security, tt.args.time); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("GoodDayToRobBank() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GoodDayToRobBank1(tt.args.security, tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoodDayToRobBank() = %v, want %v", got, tt.want)
			}
		})
	}
}
