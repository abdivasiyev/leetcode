package rotateimage

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			name: "Test #1",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "Test #2",
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rotate(tc.matrix)

			if !reflect.DeepEqual(tc.matrix, tc.expected) {
				t.Fatalf("expected: %v, got: %v\n", tc.expected, tc.matrix)
			}
		})
	}
}
