package spiralmatrix

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	testCases := []struct {
		name     string
		matrix   [][]int
		expected []int
	}{
		{
			name: "Test #1",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "Test #2",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "Test #3",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			expected: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := spiralOrder(tc.matrix)

			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("expected: %v, got: %v\n", tc.expected, got)
			}
		})
	}

}
