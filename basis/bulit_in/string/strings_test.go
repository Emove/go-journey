package string

import "testing"

func TestUpperCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{
				str: "Hello",
			},
			want: "Hello",
		},
		{
			name: "second",
			args: args{
				str: "hello",
			},
			want: "Hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperCamelCase(tt.args.str); got != tt.want {
				t.Errorf("UpperCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCamelCae(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{
				str: "hello",
			},
			want: "hello",
		},
		{
			name: "second",
			args: args{
				str: "Hello",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerCamelCae(tt.args.str); got != tt.want {
				t.Errorf("LowerCamelCae() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamelToSnake(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{
				value: "hello",
			},
			want: "hello",
		},
		{
			name: "second",
			args: args{
				value: "helloWorld",
			},
			want: "hello_world",
		},
		{
			name: "third",
			args: args{
				value: "HelloWorld",
			},
			want: "hello_world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelToSnake(tt.args.value); got != tt.want {
				t.Errorf("CamelToSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	type args struct {
		s  string
		ss string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{
				s:  "`hello`",
				ss: "`",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.args.s, tt.args.ss); got != tt.want {
				t.Errorf("Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}
