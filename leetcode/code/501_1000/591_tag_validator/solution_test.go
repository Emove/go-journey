package _91_tag_validator

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{code: "<DIV>This is the first line <![CDATA[<div>]]></DIV>"},
			want: true,
		},
		{
			name: "second",
			args: args{code: "<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>"},
			want: true,
		},
		{
			name: "third",
			args: args{code: "<A>  <B> </A>   </B>"},
			want: false,
		},
		{
			name: "forth",
			args: args{code: "<DIV>  div tag is not closed  <DIV>"},
			want: false,
		},
		{
			name: "fifth",
			args: args{code: "<DIV>  unmatched <  </DIV>"},
			want: false,
		},
		{
			name: "sixth",
			args: args{code: "<DIV> closed tags with invalid tag name  <b>123</b> </DIV>"},
			want: false,
		},
		{
			name: "seventh",
			args: args{code: "<DIV> unmatched tags with invalid tag name  </1234567890> and <CDATA[[]]>  </DIV>"},
			want: false,
		},
		{
			name: "eighth",
			args: args{code: "<DIV>  unmatched start tag <B>  and unmatched end tag </C>  </DIV>"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.code); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
