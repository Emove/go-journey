package memento

import "testing"

func TestGame(t *testing.T) {
	game := &Game{
		hp: 10,
		mp: 10,
	}

	game.status()

	progress := game.Save()

	game.Play(-2, -3)
	game.status()

	game.Load(progress)
	game.status()
}
