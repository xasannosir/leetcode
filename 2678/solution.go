package solution

import "strconv"

func countSeniors(details []string) int {
	count := 0

	for _, detail := range details {
		intVal, _ := strconv.Atoi(detail[11:13])

		if intVal > 60 {
			count++
		}
	}

	return count
}
