package _823_find_the_winner_of_the_circular_game

func FindTheWinner(n, k int) int {
	out := make([]bool, n)
	cur := 0
	for remain := n; remain > 1; remain-- {
		for i := 0; i < k-1; i++ {
			cur++
			for out[cur%n] {
				cur++
			}
		}
		out[cur%n] = true
		for out[cur%n] {
			cur++
		}
	}
	return cur%n + 1
}

func FindTheWinner1(n, k int) int {
	winner := 1
	for i := 2; i <= n; i++ {
		winner = (winner+k-1)%i + 1
	}
	return winner
}
