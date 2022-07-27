package paired69

import "testing"

func TestPaired69(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{S: "66"},
			want: "6699",
		},
		{
			name: "second",
			args: args{S: ""},
			want: "",
		},
		{
			name: "third",
			args: args{S: "9"},
			want: "69",
		},
		{
			name: "forth",
			args: args{S: "669"},
			want: "6699",
		},
		{
			name: "fifth",
			args: args{S: "696"},
			want: "6969",
		},
		{
			name: "sixth",
			args: args{S: "699"},
			want: "6699",
		},
		{
			name: "seventh",
			args: args{S: "99"},
			want: "6699",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Paired69(tt.args.S); got != tt.want {
				t.Errorf("Paired69() = %v, want %v", got, tt.want)
			}
		})
	}
}
