package tile

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTileAlg(t *testing.T) {
	type args struct {
		positions []*Position
	}
	tests := []struct {
		name string
		args args
		want *Position
	}{
		// TODO: Add test cases.
		{
			name: "first node dose not have box",
			args: args{
				positions: []*Position{
					{
						X: 2,
						Y: 1,
						Z: 1,
					},
				},
			},
			want: &Position{
				X: 1,
				Y: 1,
				Z: 1,
			},
		},
		{
			name: "insert between",
			args: args{
				positions: []*Position{
					{
						X: 1,
						Y: 1,
						Z: 1,
					},
					{
						X: 3,
						Y: 1,
						Z: 1,
					},
				},
			},
			want: &Position{
				X: 2,
				Y: 1,
				Z: 1,
			},
		},
		{
			name: "append behind",
			args: args{
				positions: []*Position{
					{
						X: 1,
						Y: 1,
						Z: 1,
					},
					{
						X: 2,
						Y: 1,
						Z: 1,
					},
				},
			},
			want: &Position{
				X: 3,
				Y: 1,
				Z: 1,
			},
		},
		{
			name: "append to last",
			args: args{
				positions: []*Position{
					{
						X: 1,
						Y: 1,
						Z: 1,
					},
					{
						X: 2,
						Y: 1,
						Z: 1,
					},
					{
						X: 3,
						Y: 1,
						Z: 1,
					},
					{
						X: 4,
						Y: 1,
						Z: 1,
					},
				},
			},
			want: &Position{
				X: 5,
				Y: 1,
				Z: 1,
			},
		},
		{
			name: "append to next layer",
			args: args{
				positions: []*Position{
					{
						X: 1,
						Y: 1,
						Z: 1,
					},
					{
						X: 2,
						Y: 1,
						Z: 1,
					},
					{
						X: 3,
						Y: 1,
						Z: 1,
					},
					{
						X: 4,
						Y: 1,
						Z: 1,
					},
					{
						X: 5,
						Y: 1,
						Z: 1,
					},
				},
			},
			want: &Position{
				X: 1,
				Y: 1,
				Z: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TileAlg(tt.args.positions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TileAlg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiTileAlg(t *testing.T) {
	// Nil Test
	//poss1 := MultiTileAlg(nil, 7)
	//for _, position := range poss1 {
	//	fmt.Println(position)
	//}
	// pre hold positionTest
	poss := make([]*Position, 0)
	poss = append(poss, &Position{X: 2, Y: 1, Z: 1})
	poss = append(poss, &Position{X: 3, Y: 1, Z: 1})
	poss = append(poss, &Position{X: 4, Y: 1, Z: 1})
	poss = append(poss, &Position{X: 2, Y: 1, Z: 2})
	poss = append(poss, &Position{X: 5, Y: 1, Z: 2})
	poss2 := MultiTileAlg(poss, 7)
	for _, position := range poss2 {
		fmt.Println(position)
	}
}
