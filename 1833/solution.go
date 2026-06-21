package solution

import "sort"

func maxIceCream(costs []int, coins int) int {
	count := 0

	sort.Ints(costs)

	for _, cost := range costs {
		if coins >= cost {
			count++
			coins -= cost
		}
	}

	return count
}
