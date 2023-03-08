package _03_kth_largest_element_in_a_stream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	kl := Constructor(3, []int{4, 5, 8, 2})
	assert.Equal(t, 4, kl.Add(3))
	assert.Equal(t, 5, kl.Add(5))
	assert.Equal(t, 5, kl.Add(10))
	assert.Equal(t, 8, kl.Add(9))
	assert.Equal(t, 8, kl.Add(4))
}
