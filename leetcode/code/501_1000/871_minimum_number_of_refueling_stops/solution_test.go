package _71_minimum_number_of_refueling_stops

import "testing"

func TestMinRefuelStops(t *testing.T) {
	type args struct {
		target    int
		startFuel int
		stations  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{target: 1, startFuel: 1, stations: [][]int{}},
			want: 0,
		},
		{
			name: "second",
			args: args{target: 100, startFuel: 1, stations: [][]int{{10, 100}}},
			want: -1,
		},
		{
			name: "third",
			args: args{target: 100, startFuel: 10, stations: [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinRefuelStopsGreedy(tt.args.target, tt.args.startFuel, tt.args.stations); got != tt.want {
				t.Errorf("MinRefuelStops() = %v, want %v", got, tt.want)
			}
		})
	}
}
