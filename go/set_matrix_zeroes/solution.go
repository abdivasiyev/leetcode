package setmatrixzeroes

func setZeroes(matrix [][]int) {
	marker := -(1 << 31) - 1

	n, m := len(matrix), len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] != 0 {
				continue
			}

			for k := 0; k < n; k++ {
				if matrix[k][j] == 0 {
					continue
				}
				matrix[k][j] = marker
			}

			for k := 0; k < m; k++ {
				if matrix[i][k] == 0 {
					continue
				}
				matrix[i][k] = marker
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] != marker {
				continue
			}
			matrix[i][j] = 0
		}
	}
}
