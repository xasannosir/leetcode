package solution

func calPoints(operations []string) int {
	var records []int
	total := 0

	for i := 0; i < len(operations); i++ {
		if operations[i] == "D" && len(records) > 0 {
			records = append(records, records[len(records)-1]*2)
		} else if operations[i] == "C" && len(records) > 0 {
			records = records[:len(records)-1]
		} else if operations[i] == "+" && len(records) > 1 {
			records = append(records, records[len(records)-1]+records[len(records)-2])
		} else {
			isNegative := false
			num := 0

			if operations[i][0] == '-' {
				isNegative = true
				operations[i] = operations[i][1:]
			}

			bytes := []byte(operations[i])

			for i := 0; i < len(bytes); i++ {
				num = num*10 + int(bytes[i]-48)
			}

			if isNegative {
				records = append(records, -1*num)
			} else {
				records = append(records, num)
			}
		}
	}

	for _, num := range records {
		total += num
	}

	return total
}
