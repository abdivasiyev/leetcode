package binary_gap

// 868. Binary Gap
// https://leetcode.com/problems/binary-gap/
//
// Given a positive integer n,
// find and return the longest distance between any two adjacent 1's
// in the binary representation of n.
// If there are no two adjacent 1's, return 0.
//
// Two 1's are adjacent if there are only 0's separating them
// (possibly no 0's). The distance between two 1's is the absolute
// difference between their bit positions.
// For example, the two 1's in "1001" have a distance of 3.

func binaryGap(n int) int {
	var (
		longestDistance = 0
		currentIndex    = 0
		indexes         [32]int
		m, i            int
	)

	// 22 -> 01101
	for n > 0 {
		m = n & 1
		if m == 1 {
			indexes[i] = currentIndex
			i++
		}
		currentIndex++
		n >>= 1
	}

	for i = 0; i < len(indexes)-1; i++ {
		longestDistance = max(longestDistance, indexes[i+1]-indexes[i])
	}

	return longestDistance
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
