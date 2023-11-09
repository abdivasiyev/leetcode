package canplaceflowers

import "testing"

func Test_canPlaceFlowers(t *testing.T) {

	testCases := []struct {
		name      string
		flowerbed []int
		n         int
		expected  bool
	}{
		{
			name:      "Test #1",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			name:      "Test #2",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			name:      "Test #3",
			flowerbed: []int{0},
			n:         1,
			expected:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := canPlaceFlowers(tc.flowerbed, tc.n)

			if got != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}
