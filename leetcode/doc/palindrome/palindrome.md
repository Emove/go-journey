## 回文数
### 1、求回文数
在求回文数时，有些时候难免要枚举。但是枚举时应从回文数本身出发，通过构造范围内回文数的方式来找到答案。
如[leetcode.479](https://leetcode-cn.com/problems/largest-palindrome-product/)，在寻找最大的回文数时，通过左半部分来构造一个回文数，这种方式，比枚举乘积后再判断是否是回文数高效的多。

```go
palindrome := left
for	x := left; x > 0; x/=10 {
	p = p * 10 + x % 10
}
```

### 2、判断一个数是否是回文数

1. 第一种办法是直接将回文数转成字符串，再用双指针进行匹配。

   ```go
   func isPalindrome(num int) bool {
   	itoa := strconv.Itoa(num)
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
   ```

2. 第二种针对与数字，将数字进行倒转，如果倒转后等于原来的数，即为回文数

   ```go
   func isPalindrome(n int) bool {
   	invert := 0
   	for x:=n; x > 0; x/=10{
   		invert=invert*10+x%10
   	}
   	return invert == n
   }
   ```

   