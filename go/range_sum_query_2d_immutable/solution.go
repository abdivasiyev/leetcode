package rangesumquery2dimmutable

import "fmt"

type NumMatrix struct {
	sumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n := len(matrix)
	if n == 0 {
		return NumMatrix{}
	}
	m := len(matrix[0])
	if m == 0 {
		return NumMatrix{}
	}

	sumMatrix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		sumMatrix[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			sumMatrix[i][j] = sumMatrix[i-1][j] + sumMatrix[i][j-1] + matrix[i-1][j-1] - sumMatrix[i-1][j-1]
		}
	}

	return NumMatrix{
		sumMatrix: sumMatrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// Sum(ABCD) = Sum(OD) - Sum(OC) - Sum(OB) + Sum(OA)
	sumOD := this.sumMatrix[row2+1][col2+1]
	sumOC := this.sumMatrix[row2+1][col1]
	sumOB := this.sumMatrix[row1][col2+1]
	sumOA := this.sumMatrix[row1][col1]

	return sumOD - sumOC - sumOB + sumOA
}

func (this *NumMatrix) Print() {
	fmt.Println("Sum matrix")
	for i := 0; i < len(this.sumMatrix); i++ {
		for j := 0; j < len(this.sumMatrix[i]); j++ {
			fmt.Printf("%d ", this.sumMatrix[i][j])
		}
		fmt.Println()
	}
}
