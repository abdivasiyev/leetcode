package toeplitzmatrix

import "testing"

func Test_isToeplitzMatrix(t *testing.T) {

	testCases := []struct {
		name     string
		matrix   [][]int
		expected bool
	}{
		{
			name: "Test #1",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 1, 2, 3},
				{9, 5, 1, 2},
			},
			expected: true,
		},
		{
			name: "Test #2",
			matrix: [][]int{
				{1, 2},
				{2, 2},
			},
			expected: false,
		},
		{
			name: "Test #3",
			matrix: [][]int{
				{18},
				{66},
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isToeplitzMatrix(tc.matrix)
			if got != tc.expected {
				t.Fatalf("expected: %v, got: %v\n", tc.expected, got)
			}
		})
	}
}
