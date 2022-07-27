package zipzap_conversion_06

import "testing"

func TestConvert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{s: "PAYPALISHIRING", numRows: 3},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "second",
			args: args{s: "PAYPALISHIRING", numRows: 4},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "third",
			args: args{s: "A", numRows: 1},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkConvert(b *testing.B) {
	Convert("PAYPALISHIRING", 3)
	Convert("PAYPALISHIRING", 4)
}
