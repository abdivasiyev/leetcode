package setmatrixzeroes

import (
	"reflect"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	testCases := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			name: "Test #1",
			matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "Test #2",
			matrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			setZeroes(tc.matrix)

			if !reflect.DeepEqual(tc.matrix, tc.expected) {
				t.Fatalf("expected: %v, got: %v\n", tc.expected, tc.matrix)
			}
		})
	}
}
