package solution

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	ret := make([][]int, 0)
	pairs := make(map[int]int)

	for _, item := range items1 {
		pairs[item[0]] += item[1]
	}

	for _, item := range items2 {
		pairs[item[0]] += item[1]
	}

	for k, v := range pairs {
		ret = append(ret, []int{k, v})
	}

	for i := 0; i < len(ret); i++ {
		for j := 0; j < len(ret)-i-1; j++ {
			if ret[j][0] > ret[j+1][0] {
				ret[j], ret[j+1] = ret[j+1], ret[j]
			}
		}
	}

	return ret
}
