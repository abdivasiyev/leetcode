package searchinsertposition

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	m := (l + r) / 2

	for l <= r {
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
		m = (l + r) / 2
	}

	return l
}
