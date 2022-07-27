package decorator

import (
	"container/list"
	"fmt"
)

type Cache interface {
	put(key, value string)
	get(key string) string
	remove(key string)
}

type PerpetualCache struct {
	maps map[string]string
}

const MaxLength = 4

type FifoCache struct {
	decorator Cache
	order     *list.List
}

func NewCache() Cache {
	return &FifoCache{
		decorator: &PerpetualCache{
			maps: make(map[string]string),
		},
		order: list.New(),
	}
}

func (cache *PerpetualCache) put(key, value string) {
	cache.maps[key] = value
}

func (cache *PerpetualCache) get(key string) string {
	return cache.maps[key]
}

func (cache *PerpetualCache) remove(key string) {
	delete(cache.maps, key)
}

func (cache *FifoCache) put(key, value string) {
	if cache.order.Len() >= MaxLength {
		front := cache.order.Front()
		cache.order.Remove(front)
		cache.decorator.remove(fmt.Sprint(front.Value))
		fmt.Printf("remove %#v\n", front)
	}
	cache.decorator.put(key, value)
	cache.order.PushBack(key)
	fmt.Println()
	for e := cache.order.Front(); e != nil; e = e.Next() {
		fmt.Printf("value : %#v\n", e.Value)
	}
}

func (cache *FifoCache) get(key string) string {
	return cache.decorator.get(key)
}

func (cache *FifoCache) remove(key string) {
	cache.decorator.remove(key)
	element := findElementByKey(cache.order, key)
	if nil != element {
		cache.order.Remove(element)
	}
}

func findElementByKey(list *list.List, key string) *list.Element {
	for e := list.Front(); e != nil; e = e.Next() {
		elementValue := fmt.Sprint(e.Value)
		if key == elementValue {
			return e
		}
	}
	return nil
}
