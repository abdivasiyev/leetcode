package canplaceflowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)
	for i := 0; i < size; i++ {
		if n == 0 {
			return true
		}
		if (i == 0 || i > 0 && flowerbed[i-1] == 0) &&
			flowerbed[i] == 0 &&
			(i == size-1 || i < size-1 && flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}

	return n == 0
}
