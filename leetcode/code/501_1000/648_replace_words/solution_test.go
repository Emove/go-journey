package _48_replace_words

import "testing"

func TestReplaceWords(t *testing.T) {
	type args struct {
		dictionary []string
		sentence   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{dictionary: []string{"cat", "bat", "rat"}, sentence: "the cattle was rattled by the battery"},
			want: "the cat was rat by the bat",
		},
		{
			name: "second",
			args: args{dictionary: []string{"a", "b", "c"}, sentence: "aadsfasf absbs bbab cadsfafs"},
			want: "a a b c",
		},
		{
			name: "third",
			args: args{dictionary: []string{"a", "aa", "aaa", "aaaa"}, sentence: "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"},
			want: "a a a a a a a a bbb baba a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceWords(tt.args.dictionary, tt.args.sentence); got != tt.want {
				t.Errorf("ReplaceWords() = %v, want= %v", got, tt.want)
			}
		})
	}
}
