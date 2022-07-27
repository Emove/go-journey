package _175_prime_arrangements

const MOD = 1e9 + 7

var PRIMES [101]int
var FACTORIAL [101]int64

func init() {
	PRIMES[0], PRIMES[1] = 0, 0
	FACTORIAL[0], FACTORIAL[1] = 1, 1
	for i := 2; i <= 100; i++ {
		FACTORIAL[i] = FACTORIAL[i-1] * int64(i) % MOD
		if isPrime(i) {
			PRIMES[i] = PRIMES[i-1] + 1
		} else {
			PRIMES[i] = PRIMES[i-1] + 0
		}
	}
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func NumPrimeArrangements(n int) int {
	return int(FACTORIAL[PRIMES[n]] * FACTORIAL[n-PRIMES[n]] % MOD)
}
