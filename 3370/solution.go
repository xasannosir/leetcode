package solution

func smallestNumber(n int) int {
	var isSetBits func(num int) bool

	isSetBits = func(num int) bool {
		if num == 0 {
			return true
		} else if num%2 == 0 {
			return false
		} else {
			return isSetBits(num / 2)
		}
	}

	for {
		if isSetBits(n) {
			return n
		}

		n++
	}
}
