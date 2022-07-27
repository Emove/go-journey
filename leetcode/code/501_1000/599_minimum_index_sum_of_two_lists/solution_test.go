package _99_minimum_index_sum_of_two_lists

import (
	"reflect"
	"testing"
)

func TestFindRestaurant(t *testing.T) {
	type args struct {
		list1 []string
		list2 []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name:    "first",
			args:    args{list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, list2: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}},
			wantAns: []string{"Shogun"},
		},
		{
			name:    "second",
			args:    args{list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, list2: []string{"KFC", "Shogun", "Burger King"}},
			wantAns: []string{"Shogun"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := FindRestaurant(tt.args.list1, tt.args.list2); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("FindRestaurant() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
