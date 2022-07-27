package _2_generate_parentheses

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name:    "first",
			args:    args{n: 1},
			wantAns: []string{"()"},
		},
		{
			name:    "second",
			args:    args{n: 3},
			wantAns: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := GenerateParenthesis(tt.args.n); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("GenerateParenthesis() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
