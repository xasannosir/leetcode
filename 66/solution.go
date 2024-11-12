package solution

func plusOne(digits []int) []int {
	isAdd := false
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			isAdd = true
		} else {
			digits[i] = digits[i] + 1
			isAdd = false
			break
		}
	}
	if isAdd {
		digits = append([]int{1}, digits...)
	}

	return digits
}
