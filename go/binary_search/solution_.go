package binarysearch

func search(nums []int, target int) int {
	// given numbers are sorted in ascending order
	// so we can use binary search

	// define 2 pointers: left and right
	l, r := 0, len(nums)-1

	// loop through while left pointer less or equal than right pointer
	for l <= r {
		// get middle of the pointers
		m := (l + r) / 2

		switch {
		// if middle of the pointers less than target increase left pointer
		case nums[m] < target:
			l = m + 1

		// if middle of the pointers greater than target decrease right pointer
		case nums[m] > target:
			r = m - 1

		// if middle of the pointers equals to target return m
		default:
			return m
		}
	}

	// if loop ends, we could not find target from list
	// so return -1
	return -1
}
