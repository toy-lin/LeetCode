package atoi

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	assert.Equal(t, 42, myAtoi("42"))
	assert.Equal(t, 42, myAtoi("   42"))
	assert.Equal(t, -44, myAtoi("-44"))
	assert.Equal(t, -44, myAtoi("  -44"))
	assert.Equal(t, 0, myAtoi("w234"))
	assert.Equal(t, 1241, myAtoi("1241 w"))
	assert.Equal(t, math.MinInt32, myAtoi("-91283472332"))
	assert.Equal(t, math.MinInt32, myAtoi("-2147483649"))
	assert.Equal(t, 1, myAtoi("+1"))
	assert.Equal(t, 0, myAtoi(""))
	assert.Equal(t, 4500, myAtoi("     +004500"))
}
