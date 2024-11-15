package solution

func mySqrt(x int) int {
	i := 1
	for ; i*i <= x; i++ {
	}

	return i - 1
}
