package remove_element

import (
	"sort"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name         string
		args         args
		expectedK    int
		expectedNums []int
	}{
		{
			name: "Test #1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			expectedK:    2,
			expectedNums: []int{2, 2},
		},
		{
			name: "Test #2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			expectedK:    5,
			expectedNums: []int{0, 0, 1, 3, 4, 2, 2, 2},
		},
		{
			name: "Test #3",
			args: args{
				nums: []int{1},
				val:  1,
			},
			expectedK:    0,
			expectedNums: []int{1},
		},
		{
			name: "Test #4",
			args: args{
				nums: []int{3, 3},
				val:  3,
			},
			expectedK:    0,
			expectedNums: []int{3, 3},
		},
		{
			name: "Test #5",
			args: args{
				nums: []int{4, 5},
				val:  5,
			},
			expectedK:    1,
			expectedNums: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK := removeElement(tt.args.nums, tt.args.val)
			if gotK != tt.expectedK {
				t.Fatalf("expected: %v, got: %v\n", tt.expectedK, gotK)
			}

			sl := tt.args.nums[0:tt.expectedK]
			sort.Slice(sl, func(i, j int) bool {
				return sl[i] < sl[j]
			})

			for i := 0; i < len(tt.expectedNums); i++ {
				if tt.expectedNums[i] != tt.args.nums[i] {
					t.Fatalf("expected: %v, got: %v\n", tt.expectedNums, tt.args.nums[:len(tt.expectedNums)])
				}
			}
		})
	}
}
