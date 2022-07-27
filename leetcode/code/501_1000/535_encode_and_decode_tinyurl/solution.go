package _35_encode_and_decode_tinyurl

import (
	"strconv"
)

type Codec struct {
	store       map[int]string
	idGenerator int
}

func Constructor() Codec {
	return Codec{store: map[int]string{}, idGenerator: 0}
}

func (c *Codec) encode(longUrl string) string {
	c.idGenerator++
	c.store[c.idGenerator] = longUrl
	return "http://tinyurl.com/" + strconv.Itoa(c.idGenerator)
}

func (c *Codec) decode(shortUrl string) string {
	i := len(shortUrl) - 1
	for ; i >= 0 && shortUrl[i] != '/'; i-- {
	}
	id, _ := strconv.Atoi(shortUrl[i+1:])
	return c.store[id]
}
