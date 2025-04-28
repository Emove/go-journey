package color

import "github.com/lucasb-eyer/go-colorful"

func Random() string {
	return colorful.HappyColor().Hex()
}
