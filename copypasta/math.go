package copypasta

func mathCollection() {
	calcGCD := func(a, b int64) int64 {
		for b > 0 {
			a, b = b, a%b
		}
		return a
	}
	calcLCM := func(a, b int64) int64 {
		return a / calcGCD(a, b) * b
	}

	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	// ax ≡ 1 (mod m)
	modInverse := func(a, m int64) int64 {
		_, x, _ := exgcd(a, m)
		return (x%m + m) % m
	}

	_ = []interface{}{calcLCM, isPrime, modInverse}
}

// exgcd solve equation ax+by=gcd(a,b)
// we have |x|<=b and |y|<=a in result (x,y)
func exgcd(a, b int64) (gcd, x, y int64) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, y, x = exgcd(b, a%b)
	y -= a / b * x
	return
}
