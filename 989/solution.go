package solution

func addToArrayForm(num []int, k int) []int {
	for i := len(num) - 1; i >= 0; i-- {
		number := num[i] + (k % 10)
		num[i] = number % 10
		k = k/10 + number/10
	}

	for k > 0 {
		num = append(append([]int{}, k%10), num...)
		k /= 10
	}

	return num
}
