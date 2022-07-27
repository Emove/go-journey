package _1_average_of_the_window

type MovingAverage struct {
	size int
	cnt  int
	sum  int
	nums []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		nums: make([]int, size, size),
	}
}

func (ma *MovingAverage) Next(val int) float64 {
	ma.sum += val
	ma.cnt++
	if ma.cnt >= ma.size {
		ma.sum -= ma.nums[ma.cnt%ma.size]
	}
	ma.nums[ma.cnt%ma.size] = val
	if ma.cnt < ma.size {
		return float64(ma.sum) / float64(ma.cnt)
	}
	return float64(ma.sum) / float64(ma.size)
}
