package searchinsertposition

import "testing"

func Test_searchInsert(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Test #1",
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "Test #2",
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "Test #3",
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			name:     "Test #4",
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := searchInsert(tc.nums, tc.target)

			if got != tc.expected {
				t.Fatalf("expected: %v, got: %v\n", tc.expected, got)
			}
		})
	}
}
