package rangesumquery2dimmutable

import "testing"

func Test_SumRegion(t *testing.T) {
	testCases := []struct {
		name     string
		matrix   [][]int
		regions  [][]int
		expected []int
	}{
		{
			name: "Test #1",
			matrix: [][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			},
			regions: [][]int{
				{2, 1, 4, 3},
				{1, 1, 2, 2},
				{1, 2, 2, 4},
			},
			expected: []int{8, 11, 12},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nm := Constructor(tc.matrix)

			nm.Print()

			for i, region := range tc.regions {
				got := nm.SumRegion(region[0], region[1], region[2], region[3])

				if got != tc.expected[i] {
					t.Fatalf("expected: %v, got: %v\n", tc.expected[i], got)
				}
			}
		})
	}
}
