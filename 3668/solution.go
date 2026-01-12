package solution

func recoverOrder(order []int, friends []int) []int {
	counter := make(map[int]struct{})
	res := make([]int, 0)

	for _, friend := range friends {
		counter[friend] = struct{}{}
	}

	for _, item := range order {
		if _, ok := counter[item]; ok {
			res = append(res, item)
		}
	}

	return res
}
