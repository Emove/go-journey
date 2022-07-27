package _44_delete_columns_to_make_sorted

import "testing"

func TestMinDeletionSize(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{strs: []string{"cba", "daf", "ghi"}},
			want: 1,
		},
		{
			name: "second",
			args: args{strs: []string{"a", "b"}},
			want: 0,
		},
		{
			name: "third",
			args: args{strs: []string{"zyx", "wvu", "tsr"}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDeletionSize(tt.args.strs); got != tt.want {
				t.Errorf("MinDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
