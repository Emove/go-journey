package _37_reorder_data_in_log_files

import (
	"sort"
	"strings"
	"unicode"
)

type Log struct {
	t       bool
	sign    string
	content string
}

func ReorderLogFiles1(logs []string) []string {
	sort.Slice(logs, func(i, j int) bool {
		l1, l2 := logs[i], logs[j]
		s1 := strings.SplitN(l1, " ", 2)[1]
		s2 := strings.SplitN(l2, " ", 2)[1]
		isDigit1 := unicode.IsDigit(rune(s1[0]))
		isDigit2 := unicode.IsDigit(rune(s2[0]))
		if isDigit1 && isDigit2 {
			return false
		}
		if !isDigit1 && !isDigit2 {
			return s1 < s2 || s1 == s2 && l1 < l2
		}
		return !isDigit1
	})
	return logs
}

func ReorderLogFiles(logs []string) []string {
	//logSlice := make([]*Log, len(logs))
	logM := make(map[string]*Log, len(logs))
	sort.SliceStable(logs, func(i, j int) bool {
		I := logs[i]
		if _, ok := logM[I]; !ok {
			logM[I] = getSign(I)
		}
		J := logs[j]
		if _, ok := logM[J]; !ok {
			logM[J] = getSign(J)
		}
		log1, log2 := logM[I], logM[J]
		if log1.t && log2.t {
			return false
		}
		if !log1.t && !log2.t {
			return log1.content < log2.content || log1.content == log2.content && log1.sign < log2.sign
		}
		return !log1.t
	})
	return logs
}

func getSign(log string) *Log {
	k := 0
	for ; k < len(log) && log[k] != ' '; k++ {
	}
	content := log[k+1:]
	return &Log{
		t:       unicode.IsDigit(rune(content[0])),
		sign:    log[:k],
		content: content,
	}
}
