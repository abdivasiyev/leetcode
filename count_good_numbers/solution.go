package count_good_numbers

const M = 1_000_000_007

func binpow(n, m int64) int64 {
	var res int64 = 1

	for m > 0 {
		if m&1 == 1 {
			res = res * n % M
		}
		n = n * n % M
		m >>= 1
	}

	return res
}

func countGoodNumbers(n int64) int {
	if n&1 == 1 {
		return int(binpow(5, n/2+1) * binpow(4, n/2) % M)
	}
	return int(binpow(5, n/2) * binpow(4, n/2) % M)
}
