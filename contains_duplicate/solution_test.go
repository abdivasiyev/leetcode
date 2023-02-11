package containsduplicate

import "testing"

func Test_containsDuplicate(t *testing.T) {

	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Test #1",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "Test #2",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
	}

	t.Parallel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := containsDuplicate(tc.nums)

			if got != tc.expected {
				t.Fatalf("expected: %v, got: %v\n", tc.expected, got)
			}
		})
	}
}
