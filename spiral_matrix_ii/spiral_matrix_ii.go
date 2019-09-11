package spiral_matrix_ii

///给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
//
//示例:
//
//输入: 3
//输出:
//[
// [ 1, 2, 3 ],
// [ 8, 9, 4 ],
// [ 7, 6, 5 ]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/spiral-matrix-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func generateMatrix(n int) (result [][]int) {
	if n == 0 {
		return
	}
	result = make([][]int, n)
	for r := range result {
		result[r] = make([]int, n)
	}

	var rowTop, rowBot, colLeft, colRight, i int
	rowBot = n - 1
	colRight = n - 1
	var ind = 1
	for {
		for i = colLeft; i < colRight+1; i++ {
			result[rowTop][i] = ind
			ind++
		}
		rowTop++
		if rowTop > rowBot {
			return
		}

		for i = rowTop; i < rowBot+1; i++ {
			result[i][colRight] = ind
			ind++
		}
		colRight--
		if colRight < colLeft {
			return
		}

		for i = colRight; i >= colLeft; i-- {
			result[rowBot][i] = ind
			ind++
		}
		rowBot--
		if rowBot < rowTop {
			return
		}

		for i = rowBot; i >= rowTop; i-- {
			result[i][colLeft] = ind
			ind++
		}
		colLeft++
		if colLeft > colRight {
			return
		}
	}
}
