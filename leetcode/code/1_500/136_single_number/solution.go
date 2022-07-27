package _36_single_number

func SingleNumber(nums []int) int {
	n := len(nums)
	cnt := make(map[int]int, n)
	for _, num := range nums {
		if v, ok := cnt[num]; ok {
			cnt[num] = v + 1
		} else {
			cnt[num] = 1
		}
	}
	for k, v := range cnt {
		if v == 1 {
			return k
		}
	}
	return -1
}

func SingleNumberByXOR(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
