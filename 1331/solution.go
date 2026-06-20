package solution

import "sort"

func arrayRankTransform(arr []int) []int {
	start := 1
	n := len(arr)
	temp := make([]int, n)
	res := make([]int, n)
	order := make(map[int]int)

	copy(temp, arr)

	sort.Ints(temp)

	for _, num := range temp {
		if _, ok := order[num]; !ok {
			order[num] = start
			start++
		}
	}

	for i, num := range arr {
		res[i] = order[num]
	}

	return res
}
