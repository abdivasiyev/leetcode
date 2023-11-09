package validanagram

import "sort"

func isAnagram(s, t string) bool {
	var (
		rS = []rune(s)
		rT = []rune(t)
	)

	sort.Slice(rS, func(i, j int) bool {
		return rS[i] < rS[j]
	})

	sort.Slice(rT, func(i, j int) bool {
		return rT[i] < rT[j]
	})

	return string(rS) == string(rT)
}
