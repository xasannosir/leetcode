package solution

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var sum int

	for i := 0; i < len(hours); i++ {
		if hours[i] >= target {
			sum++
		}
	}

	return sum
}
