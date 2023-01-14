package maxavgsubarray

import (
	"math"
	"testing"
)

func Test_findMaxAverage(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected float64
	}{
		{
			name:     "Test #1",
			nums:     []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75000,
		},
		{
			name:     "Test #2",
			nums:     []int{5},
			k:        1,
			expected: 5.00000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := findMaxAverage(tc.nums, tc.k)

			if math.Abs(got-tc.expected) > 0.00001 {
				t.Fatalf("expected: %v, got: %v\n", tc.expected, got)
			}
		})
	}
}
