package _044_count_number_of_maximum_bitwise_or_subsets

func countMaxOrSubsets(nums []int) int {
	max, ans := 0, 0
	n := len(nums)
	for i := 0; i < (1 << n); i++ {
		or := 0
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				or |= nums[j]
			}
		}
		if or > max {
			max = or
			ans = 1
		} else if or == max {
			ans++
		}
	}
	return ans
}

func dp(nums []int) int {
	n := len(nums)

	// 1 -> 0 10 -> 1 100 -> 2
	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		m[1<<i] = i
	}

	mask := 1 << n
	f := make([]int, mask)
	max, ans := 0, 0
	for state := 1; state < mask; state++ {
		// 取得最低为1 10010 -> 10
		lowOnceBit := state & -state
		// 10010 - 10 -> 10000
		prev := state - lowOnceBit
		// 10为1
		lowIndex := m[lowOnceBit]
		// 此处10000<10010 肯定已经被计算过了
		// f[10010] = f[10000] | nums[10]
		f[state] = f[prev] | nums[lowIndex]
		if f[state] > max {
			max = f[state]
			ans = 1
		} else if f[state] == max {
			ans++
		}
	}
	return ans
}

func dfs(nums []int) int {
	n := len(nums)
	max := 0
	for i := 0; i < n; i++ {
		// 先取得最大值
		max |= nums[i]
	}

	ans := 0
	var d func(pos, val int)
	d = func(pos, val int) {
		if val == max {
			// 如果当前子集的异或和已经等于max，则后续的所有子集都能符合条件
			ans += 1 << (n - pos)
			return
		}
		if pos == n {
			return
		}
		d(pos+1, val)
		d(pos+1, val|nums[pos])
	}
	d(0, 0)
	return ans
}
