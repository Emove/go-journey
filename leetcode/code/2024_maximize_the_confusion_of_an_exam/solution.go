package _024_maximize_the_confusion_of_an_exam

func MaxConsecutiveAnswers(answerKey string, k int) int {
	n := len(answerKey)
	return max(count(answerKey, k, n, 'T'), count(answerKey, k, n, 'F'))
}

func count(answerKey string, k, n int, ch byte) int {
	ans, cnt := k, 0
	for left, right := 0, 0; right < n; right++ {
		if ch != answerKey[right] {
			cnt++
		}
		if cnt > k {
			if answerKey[left] != ch {
				cnt--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
