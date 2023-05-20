package binpow

func BinPowRecursive(a, n int64) int64 {
	if n == 0 {
		return 1
	}

	res := BinPowRecursive(a, n>>1)

	if n&1 == 0 {
		return res * res
	}

	return res * res * a
}

func BinPowIterative(a, n int64) int64 {
	var res int64 = 1

	for n > 0 {
		if n&1 == 1 {
			res = res * a
		}
		a = a * a
		n >>= 1
	}

	return res
}
