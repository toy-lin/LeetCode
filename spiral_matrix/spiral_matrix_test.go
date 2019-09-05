package spiral_matrix

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	//var a = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//fmt.Println(a[0:2][2])
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
