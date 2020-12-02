package memento

import "fmt"

// 备忘录模式用于保存程序内部状态到外部，又不希望暴露内部状态的情形
// 程序内部状态使用窄接口船体给外部进行存储，从而不暴露程序实现细节
// 备忘录模式同时可以离线保存内部状态，如保存到数据库，文件等

type Memento interface {
}

type Game struct {
	hp, mp int
}

type GameMemento struct {
	hp, mp int
}

func (game *Game) Play(mpDelta, hpDelta int) {
	game.mp += mpDelta
	game.hp += hpDelta
}

func (game *Game) Save() Memento {
	return &GameMemento{
		hp: game.hp,
		mp: game.mp,
	}
}

func (game *Game) Load(memento Memento) {
	gm := memento.(*GameMemento)
	game.mp = gm.mp
	game.hp = gm.hp
}

func (game *Game) status() {
	fmt.Printf("hp: %d, mp: %d \n", game.hp, game.mp)
}
