package _33_minimum_genetic_mutation

import "testing"

func TestMinMutation(t *testing.T) {
	type args struct {
		start string
		end   string
		bank  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{start: "AACCGGTT", end: "AACCGGTA", bank: []string{"AACCGGTA"}},
			want: 1,
		},
		{
			name: "second",
			args: args{start: "AAAAACCC", end: "AACCCCCC", bank: []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMutation(tt.args.start, tt.args.end, tt.args.bank); got != tt.want {
				t.Errorf("MinMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
