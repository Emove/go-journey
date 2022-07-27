package _33_minimum_genetic_mutation

func MinMutation(start, end string, bank []string) int {
	if start == end {
		return 0
	}
	bankSet := make(map[string]struct{}, len(bank))
	for _, v := range bank {
		bankSet[v] = struct{}{}
	}
	if _, ok := bankSet[end]; !ok {
		return -1
	}
	queue := []string{start}
	for step := 0; queue != nil; step++ {
		q := queue
		queue = nil
		for _, gen := range q {
			for i, x := range gen {
				for _, ch := range "ACGT" {
					if ch == x {
						continue
					}
					g := gen[:i] + string(ch) + gen[i+1:]
					if _, ok := bankSet[g]; !ok {
						continue
					}
					if g == end {
						return step + 1
					}
					queue = append(queue, g)
					delete(bankSet, g)
				}
			}
		}
	}
	return -1
}
