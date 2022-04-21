package _24_goat_latin

import "testing"

func TestToGoatLatin(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{sentence: "I speak Goat Latin"},
			want: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{
			name: "second",
			args: args{sentence: "The quick brown fox jumped over the lazy dog"},
			want: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToGoatLatin(tt.args.sentence); got != tt.want {
				t.Errorf("ToGoatLatin() = %v,\n want = %v", got, tt.want)
			}
		})
	}
}
