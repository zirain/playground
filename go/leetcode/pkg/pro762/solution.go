package pro762

func countPrimeSetBits(left int, right int) int {
	cnt := 0
	for i := left; i <= right; i++ {
		if isPrime(countBits(i)) {
			cnt++
		}
	}

	return cnt
}

func countBits(n int) int {
	cnt := 0
	for n > 0 {
		if n&1 == 1 {
			cnt++
		}
		n = n >> 1
	}

	return cnt
}

var primes = map[int]struct{}{
	2:  {},
	3:  {},
	5:  {},
	7:  {},
	11: {},
	13: {},
	17: {},
	19: {},
}

func isPrime(n int) bool {
	_, ok := primes[n]
	return ok
}
