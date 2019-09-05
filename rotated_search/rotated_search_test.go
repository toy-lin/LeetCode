package rotated_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestF(t *testing.T) {
	assert.Equal(t, 1, search([]int{1, 3}, 3))
	assert.Equal(t, 1, search([]int{3, 1}, 1))
	assert.Equal(t, 2, search([]int{1, 2, 3}, 3))
	assert.Equal(t, 1, search([]int{1, 2, 3}, 2))
	assert.Equal(t, 0, search([]int{1, 2, 3}, 1))
	assert.Equal(t, 3, search([]int{4, 5, 6, 1, 2, 3}, 1))
	assert.Equal(t, 2, search([]int{4, 5, 6, 1, 2, 3}, 6))
	assert.Equal(t, 1, search([]int{4, 5, 6, 1, 2, 3}, 5))
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8))
	assert.Equal(t, 3, search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 7))
	assert.Equal(t, 2, search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 6))
	assert.Equal(t, 5, search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 1))
	assert.Equal(t, 6, search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 2))
}
