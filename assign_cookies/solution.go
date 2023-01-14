package assign_cookies

import "sort"

// 455. Assign Cookies
// https://leetcode.com/problems/assign-cookies/
//
// Assume you are an awesome parent and want to give
// your children some cookies.
// But, you should give each child at most one cookie.
//
// Each child i has a greed factor g[i],
// which is the minimum size of a cookie that the child will be content with;
// and each cookie j has a size s[j].
// If s[j] >= g[i], we can assign the cookie j to the child i,
// and the child i will be content.
// Your goal is to maximize the number of your content children
// and output the maximum number.

func findContentChildren(g []int, s []int) int {
	var (
		maxContent int
		i, j       = 0, 0
	)

	sort.Ints(g)
	sort.Ints(s)

	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			maxContent++
			i++
		}
		j++
	}

	return maxContent
}
