package maximumproductofthreenumbers

import "testing"

func Test_maximumProduct(t *testing.T) {

	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Test #1",
			nums:     []int{1, 2, 3},
			expected: 6,
		},
		{
			name:     "Test #2",
			nums:     []int{1, 2, 3, 4},
			expected: 24,
		},
		{
			name:     "Test #3",
			nums:     []int{-1, -2, -3},
			expected: -6,
		},
		{
			name:     "Test #4",
			nums:     []int{-100, -98, -1, 2, 3, 4},
			expected: 39200,
		},
		{
			name:     "Test #5",
			nums:     []int{-8, -7, -2, 10, 20},
			expected: 1120,
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := maximumProduct(tc.nums)

			if got != tc.expected {
				t.Fatalf("expected: %v, got: %v\n", tc.expected, got)
			}
		})
	}
}
