package solution

func isBalanced(num string) bool {
	evenSum, oddSum := 0, 0

	for i := 0; i < len(num); i++ {
		if i%2 == 0 {
			evenSum += int(byte(num[i]) - 48)
		} else {
			oddSum += int(byte(num[i]) - 48)
		}
	}

	return evenSum == oddSum
}
