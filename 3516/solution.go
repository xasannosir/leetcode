package solution

func findClosest(x int, y int, z int) int {
	abs := func(num int) int {
		if num < 0 {
			return -num
		}

		return num
	}

	if abs(x-z) < abs(y-z) {
		return 1
	} else if abs(x-z) > abs(y-z) {
		return 2
	} else {
		return 0
	}
}
