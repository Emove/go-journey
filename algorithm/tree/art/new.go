package art

import art "github.com/plar/go-adaptive-radix-tree"

var tree art.Tree

func init() {
	tree = art.New()
}

func set(key string, value interface{}) {
	tree.Insert([]byte(key), value)
}

func find(key string) (interface{}, bool) {
	return tree.Search([]byte(key))
}
