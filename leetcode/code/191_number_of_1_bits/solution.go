package _91_number_of_1_bits

import "fmt"

func HammingWeight(num uint32) (ones int) {
	for ; num > 0; num &= num - 1 {
		ones++
		fmt.Printf("num: %d, ones: %d\n", num, ones)
	}
	return
}
