// Package range_sum_query
//
// 304. Range Sum Query 2D - Immutable
// Given a 2D matrix matrix, handle multiple queries of the following type:
//
// Calculate the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1)
// and lower right corner (row2, col2).
// Implement the NumMatrix class:
//
// NumMatrix(int[][] matrix) Initializes the object with the integer matrix matrix.
// int sumRegion(int row1, int col1, int row2, int col2)
// Returns the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1)
// and lower right corner (row2, col2).
package range_sum_query

type NumMatrix struct {
	matrix [][]int
}

/**
3  0  1  4  2
5  6  3  2  1
1  2  0  1  5
4  1  0  1  7
1  0  3  0  5

0  0  0  0  0  0
0  3  3  4  8 10
0  8 14 18 24 27
0  9 17 21 28 36
0 13 22 26 34 49
0 14 23 30 38 58

(1, 1) => (0 + 0) + (3 - 0)
(1, 2) => (0 + 3) + (0 - 0)
(1, 3) => (0 + 3) + (1 - 0)
*/

func Constructor(matrix [][]int) NumMatrix {
	n := len(matrix)
	if n == 0 {
		return NumMatrix{}
	}
	m := len(matrix[0])
	nm := NumMatrix{matrix: make([][]int, n+1)}
	for i := 0; i < n+1; i++ {
		nm.matrix[i] = make([]int, m+1)
	}

	for i := range matrix {
		for j := range matrix[i] {
			nm.matrix[i+1][j+1] = nm.matrix[i+1][j] + nm.matrix[i][j+1] + matrix[i][j] - nm.matrix[i][j]
		}
	}

	return nm
}

func (nm *NumMatrix) SumRegion(y1, x1, y2, x2 int) int {
	return nm.matrix[y2+1][x2+1] - nm.matrix[y1][x2+1] - nm.matrix[y2+1][x1] + nm.matrix[y1][x1]
}
