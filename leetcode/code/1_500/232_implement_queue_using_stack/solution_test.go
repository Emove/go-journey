package _32_implement_queue_using_stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	assert.Equal(t, 1, queue.Peek())
	assert.Equal(t, 1, queue.Pop())
	assert.False(t, queue.Empty())
}
