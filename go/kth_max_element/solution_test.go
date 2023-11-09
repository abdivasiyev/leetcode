package kthmaxelement

import "testing"

func Test_findKthLargest(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Test #1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "Test #2",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := findKthLargest(tc.nums, tc.k)

			if got != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}
