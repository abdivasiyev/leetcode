package spiralmatrix

import "fmt"

func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	result := make([]int, n*m)
	resultIndex := 0
	visited := 1 << 10
	i, j := 0, 0
	direction := 0 // 0 - right, 1 - bottom, 2 - left, 3 - top

	for resultIndex < n*m {
		fmt.Println(i, j, direction)
		if matrix[i][j] != visited {
			result[resultIndex], matrix[i][j] = matrix[i][j], visited
			resultIndex++
		}

		if direction%4 == 0 { // right
			if j+1 < m && matrix[i][j+1] != visited {
				j++
			} else {
				direction = (direction + 1) % 4
			}
		} else if direction%4 == 1 { // bottom
			if i+1 < n && matrix[i+1][j] != visited {
				i++
			} else {
				direction = (direction + 1) % 4
			}
		} else if direction%4 == 2 { // left
			if j-1 >= 0 && matrix[i][j-1] != visited {
				j--
			} else {
				direction = (direction + 1) % 4
			}
		} else if direction%4 == 3 { // top
			if i-1 >= 0 && matrix[i-1][j] != visited {
				i--
			} else {
				direction = (direction + 1) % 4
			}
		}
	}
	return result
}
