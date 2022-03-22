package _038_remove_colored_pieces_if_both_neighbors_are_the_same_color

func WinnerOfGame(colors string) bool {
	cnt := 0
	for i := 0; i < len(colors)-2; i++ {
		seg := colors[i : i+3]
		if seg == "AAA" {
			cnt++
		} else if seg == "BBB" {
			cnt--
		}
	}
	return cnt > 0
}
