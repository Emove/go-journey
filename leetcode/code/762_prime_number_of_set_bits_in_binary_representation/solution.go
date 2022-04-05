package _62_prime_number_of_set_bits_in_binary_representation

import "math/bits"

// 2,3,5,7,11,13,17,19
var prime = []bool{false, false, true, true, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true}

// 从低位到高位 10100010100010101100
const MASK = 665772

func CountPrimeSetBits(left, right int) (ans int) {
	for i := left; i <= right; i++ {
		x, cnt := i, 0
		for x != 0 {
			x -= x & -x
			cnt++
		}
		if prime[cnt] {
			ans++
		}
	}
	return
}

func CountPrimeSetBits1(left, right int) (ans int) {
	for i := left; i <= right; i++ {
		if prime[bits.OnesCount(uint(i))] {
			ans++
		}
	}
	return
}

func CountPrimeSetBits2(left, right int) (ans int) {
	for x := left; x <= right; x++ {
		if 1<<bits.OnesCount(uint(x))&MASK != 0 {
			ans++
		}
	}
	return
}
