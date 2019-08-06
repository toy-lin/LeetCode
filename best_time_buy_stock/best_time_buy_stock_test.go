package best_time_buy_stock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBestTimeBuy(t *testing.T) {
	ints := []int{7, 1, 5, 3, 6, 4}
	assert.Equal(t, maxProfit(ints), 7)
	ints = []int{6, 1, 3, 2, 4, 7}
	assert.Equal(t, maxProfit(ints), 7)
}
