package _79_largest_palindrome_product

import (
	"math"
	"strconv"
)

var answer = []int{9, 987, 123, 597, 677, 1218, 877, 475}

const MOD_NUM = 1337

func LargestPalindrome(n int) int {
	if n == 1 {
		return 9
	}

	max := int(math.Pow(float64(10), float64(n)) - 1)

	for left := max; ; left-- {
		palindrome := left
		for x := left; x > 0; x /= 10 {
			palindrome = palindrome*10 + x%10
		}
		for x := max; x*x >= palindrome; x-- {
			if palindrome%x == 0 {
				return palindrome % MOD_NUM
			}
		}
	}
}

func isPalindrome(num uint64) bool {
	itoa := strconv.FormatUint(num, 10)
	left, right := 0, len(itoa)-1
	for left < right {
		if itoa[left] != itoa[right] {
			return false
		}
		left++
		right--
	}
	return true
}
