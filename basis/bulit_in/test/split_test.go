package test

import (
	"reflect"
	"testing"
)

type splitClass struct {
	input string
	sep   string
	want  []string
}

// 测试函数必须以Test开头，且必须接受一个*testing.T类型参数
func TestSplit(t *testing.T) {
	tests := []splitClass{
		{input: "a,b,c", sep: ",", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
	}

	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted: %#v, got:%#v", tc.want, got)
		}
	}
}

func TestSubtestSplit(t *testing.T) {
	tests := map[string]splitClass{
		"first":  {input: "a,b,c", sep: ",", want: []string{"a", "b", "c"}},
		"second": {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"third":  {input: "abcd", sep: "bc", want: []string{"a", "d"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted: %#v, got:%#v", tc.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a,b,c", ",")
	}
}
