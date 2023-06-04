package remove_element

func removeElement(nums []int, val int) int {
	i := 0

	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	return i
}
