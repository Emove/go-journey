package _88_longest_absolute_file_path

func LengthLongestPath(input string) int {
	ans := ""
	n := len(input)
	levels := make(map[int]string)
	for i := 0; i < n; {
		level := 0
		for ; i < n && input[i] == '\t'; i++ {
			level++
		}

		j, isDir := i, true
		for ; j < n && input[j] != '\n'; j++ {
			if input[j] == '.' {
				isDir = false
			}
		}

		// 当前文件
		cur := input[i:j]
		path := levels[level-1]
		if len(path) == 0 {
			path = cur
		} else {
			path += "/" + cur
		}

		if isDir {
			levels[level] = path
		} else if ans == "" || len(path) > len(ans) {
			ans = path
		}
		i = j + 1
	}
	return len(ans)
}
