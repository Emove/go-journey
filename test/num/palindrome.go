package num

import "fmt"

func isPalindrome(n int) bool {
	invert := 0
	for x := n; x > 0; x /= 10 {
		invert = invert*10 + x%10
	}
	fmt.Println(invert)
	return invert == n
}
