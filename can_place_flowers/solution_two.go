package canplaceflowers

func _canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if n == 0 {
			return true
		}

		switch {
		case i == 0 && i+1 < len(flowerbed)-1:
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		case i-1 >= 0 && i < len(flowerbed)-1:
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		case i-1 >= 0 && i == len(flowerbed)-1:
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 {
				flowerbed[i] = 1
				n--
			}
		case i+1 <= len(flowerbed)-1:
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		default:
			if flowerbed[i] == 0 {
				n--
			}
		}
	}

	return n == 0
}
