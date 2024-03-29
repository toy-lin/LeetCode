package spiral_matrix

//给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
//
//示例 1:
//
//输入:
//[
// [ 1, 2, 3 ],
// [ 4, 5, 6 ],
// [ 7, 8, 9 ]
//]
//输出: [1,2,3,6,9,8,7,4,5]
//示例 2:
//
//输入:
//[
//  [1, 2, 3, 4],
//  [5, 6, 7, 8],
//  [9,10,11,12]
//]
//输出: [1,2,3,4,8,12,11,10,9,5,6,7]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/spiral-matrix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func spiralOrder(matrix [][]int) (arr []int) {
	if len(matrix) == 0 {
		return
	}
	var rowTop, rowBot, colLeft, colRight, i int
	rowBot = len(matrix) - 1
	colRight = len(matrix[0]) - 1
	for {
		arr = append(arr, matrix[rowTop][colLeft:colRight+1]...)
		rowTop++
		if rowTop > rowBot {
			return
		}

		for i = rowTop; i < rowBot+1; i++ {
			arr = append(arr, matrix[i][colRight])
		}
		colRight--
		if colRight < colLeft {
			return
		}

		for i = colRight; i >= colLeft; i-- {
			arr = append(arr, matrix[rowBot][i])
		}
		rowBot--
		if rowBot < rowTop {
			return
		}

		for i = rowBot; i >= rowTop; i-- {
			arr = append(arr, matrix[i][colLeft])
		}
		colLeft++
		if colLeft > colRight {
			return
		}
	}
}
