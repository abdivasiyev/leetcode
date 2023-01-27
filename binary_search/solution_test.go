package binarysearch

import "testing"

func Test_search(t *testing.T) {
	// Run test cases parallel
	t.Parallel()

	// Test cases for binary search
	testCases := []struct {
		name     string // test name
		nums     []int  // list of numbers
		target   int    // target for finding from list
		expected int    // expected result from `search` function
	}{
		{
			name:     "Test #1",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			name:     "Test #2",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
		{
			name:     "Test #3",
			nums:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "Test #4",
			nums:     []int{1, 2, 3, 4, 10, 15},
			target:   15,
			expected: 5,
		},
	}

	for _, tc := range testCases {
		// Run sub test
		t.Run(tc.name, func(t *testing.T) {
			// execute `search` function
			got := search(tc.nums, tc.target)

			// check got value equals or not to expected value
			if tc.expected != got {
				// stop running test case if expected value is not equal to got value
				t.Fatalf("expected: %v, got: %v\n", tc.expected, got)
			}
		})
	}
}
