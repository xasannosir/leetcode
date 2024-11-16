package solution

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	} else if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	tmp := myPow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	} else {
		return x * tmp * tmp
	}
}
