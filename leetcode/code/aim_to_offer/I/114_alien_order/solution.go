package _14_alien_order

// AlienOrder
// 解法：
//	1、遍历所有字符，建立图，并额外记录入度
//	2、对图进行拓扑排序，得到的顺序即为alienOrder
// 	3、若图中存在环，则说明不存在合法字母顺序
func AlienOrder(words []string) string {
	edges, indegrees := make([][]int, 26), make(map[byte]int, 26)
	for i := range edges {
		edges[i] = make([]int, 0)
	}
	for i := 0; i < len(words[0]); i++ {
		indegrees[words[0][i]] = 0
	}
	var cover bool
	for i := 1; i < len(words); i++ {
		prev, cur := words[i-1], words[i]
		for _, ch := range cur {
			indegrees[byte(ch)] += 0
		}
		cover = true
		for j := 0; j < len(prev) && j < len(cur); j++ {
			if prev[j] != cur[j] {
				cover = false
				edges[prev[j]-'a'] = append(edges[prev[j]-'a'], int(cur[j]-'a'))
				indegrees[cur[j]]++
				break
			}
		}
		if cover && len(prev) > len(cur) {
			// prev长于cur，并且prev能够覆盖cur，不合法，很可能存在环
			// 即prev='abcd', cur='abc'
			return ""
		}
	}

	order := make([]byte, len(indegrees))
	queue := order[:0]
	for ch, in := range indegrees {
		if in == 0 {
			queue = append(queue, ch)
		}
	}

	for len(queue) > 0 {
		vert := queue[0] - 'a'
		queue = queue[1:]
		for _, v := range edges[vert] {
			indegrees[byte(v+'a')]--
			if indegrees[byte(v+'a')] == 0 {
				queue = append(queue, byte(v+'a'))
			}
		}
	}

	if cap(queue) == 0 {
		return string(order)
	}

	return ""
}
