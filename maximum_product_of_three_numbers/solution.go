package maximumproductofthreenumbers

import (
	"math"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt, math.MaxInt
	max1, max2, max3 := math.MinInt, math.MinInt, math.MinInt

	for i := 0; i < len(nums); i++ {
		if min1 > nums[i] {
			min1, min2 = nums[i], min1
		} else if min2 > nums[i] {
			min2 = nums[i]
		}

		if max1 < nums[i] {
			max1, max2, max3 = nums[i], max1, max2
		} else if max2 < nums[i] {
			max2, max3 = nums[i], max2
		} else if max3 < nums[i] {
			max3 = nums[i]
		}
	}

	return max(
		min1*min2*max1,
		max1*max2*max3,
	)
}

func maximumProduct_Solution1(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	n := len(nums)

	return max(
		nums[0]*nums[1]*nums[n-1],
		nums[n-1]*nums[n-2]*nums[n-3],
	)
}
