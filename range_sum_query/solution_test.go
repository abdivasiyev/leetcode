package range_sum_query

import "testing"

func TestSumRegion(t *testing.T) {
	testCases := []struct {
		name    string
		matrix  [][]int
		inputs  [][]int
		results []int
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
			inputs: [][]int{
				{2, 1, 4, 3},
				{1, 1, 2, 2},
				{1, 2, 2, 4},
			},
			results: []int{8, 11, 12},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nm := Constructor(tc.matrix)

			for i, input := range tc.inputs {
				result := nm.SumRegion(input[0], input[1], input[2], input[3])

				if tc.results[i] != result {
					t.Errorf("expected %v, got %v\n", tc.results[i], result)
					t.Fail()
				}
			}
		})
	}
}
