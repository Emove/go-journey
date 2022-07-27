package char_and_num_return

import (
	"sort"
	"strconv"
	"strings"
)

func char_and_num_return(text_source string) string {
	// write code here
	ans := strings.Builder{}
	nums := make([]int, 0)
	for i := 0; i < len(text_source); i++ {
		if isLetter(text_source[i]) {
			j := i
			for ; j < len(text_source) && text_source[j] != ' '; j++ {
			}
			ans.WriteString(text_source[i : j+1])
			i = j
		} else {
			num := 0
			j := i
			for ; j < len(text_source) && text_source[j] != ' '; j++ {
				num = num*10 + int(text_source[j]-'0')
			}
			i = j
			nums = append(nums, num)
		}
	}

	if len(nums) > 0 {
		sort.Ints(nums)
		for i := 0; i < len(nums); i++ {
			ans.WriteString(strconv.Itoa(nums[i]))
			ans.WriteByte(' ')
		}
	}
	if ans.Len() > 0 {
		return ans.String()[:ans.Len()-1]
	}
	return ans.String()
}

func isLetter(ch byte) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}
