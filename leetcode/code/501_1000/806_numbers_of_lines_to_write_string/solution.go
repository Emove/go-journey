package _06_numbers_of_lines_to_write_string

func NumberOfLines(widths []int, s string) []int {
	sum, lines := 0, 1
	for _, v := range s {
		if sum+widths[v-'a'] > 100 {
			lines++
			sum = 0
		}
		sum += widths[v-'a']
	}
	return []int{lines, sum}
}
