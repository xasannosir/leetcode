package solution

func finalValueAfterOperations(operations []string) int {
	var sum int

	for _, operation := range operations {
		if operation == "++X" || operation == "X++" {
			sum += 1
		} else {
			sum -= 1
		}
	}

	return sum
}
