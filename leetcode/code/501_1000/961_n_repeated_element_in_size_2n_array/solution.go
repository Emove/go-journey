package _61_n_repeated_element_in_size_2n_array

func RepeatedNTimes(nums []int) int {
	n := len(nums) / 2
	table := make([][]int, n)
	for _, num := range nums {
		index := num % n
		l := table[index]
		if len(l) == 0 {
			table[index] = []int{num}
		} else {
			for _, v := range l {
				if v == num {
					return num
				}
			}
			table[index] = append(l, num)
		}
	}
	return -1
}
