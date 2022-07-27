package _0_valid_parentheses

func IsValid(s string) bool {
	stk := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stk = append(stk, v)
		} else {
			if v == 41 {
				//ascii中，‘(’与‘)’相差一位
				v -= 1
			} else {
				// '['、'{'与']'、'}'各相差两位
				v -= 2
			}
			if len(stk) == 0 || v != stk[len(stk)-1] {
				return false
			}
			stk = stk[:len(stk)-1]
		}
	}
	return len(stk) == 0
}
