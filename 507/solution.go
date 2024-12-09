package solution

func checkPerfectNumber(num int) bool {
	var sum = 0
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	return sum == num
}
