package solution

func sumOfTheDigitsOfHarshadNumber(x int) int {
	digitsSum := 0
	temp := x

	for temp > 0 {
		digitsSum += temp % 10
		temp /= 10
	}

	if x%digitsSum == 0 {
		return digitsSum
	} else {
		return -1
	}
}
