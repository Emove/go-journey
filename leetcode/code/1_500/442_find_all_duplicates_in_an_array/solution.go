package _42_find_all_duplicates_in_an_array

func FindDuplicates(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if v < 0 || i+1 == v {
			continue
		}
		if nums[v-1] == v {
			res = append(res, v)
			nums[v-1] = -nums[v-1]
		} else {
			nums[i], nums[v-1] = nums[v-1], v
			i--
		}
	}
	return res
}

func FindDuplicates1(nums []int) (ans []int) {
	for _, x := range nums {
		if x < 0 {
			x = -x
		}
		if nums[x-1] > 0 {
			nums[x-1] = -nums[x-1]
		} else {
			ans = append(ans, x)
		}
	}
	return
}
