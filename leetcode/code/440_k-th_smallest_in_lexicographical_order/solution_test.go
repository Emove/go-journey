package _40_k_th_smallest_in_lexicographical_order

import "testing"

func TestFindKthNumber(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{n: 13, k: 2},
			want: 10,
		},
		{
			name: "second",
			args: args{n: 10, k: 3},
			want: 2,
		},
	}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := FindKthNumber(tt.args.n, tt.args.k); got != tt.want {
	//			t.Errorf("FindKthNumber() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthNumber1(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("FindKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
