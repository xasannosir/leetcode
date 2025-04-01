package solution

func duplicateZeros(arr []int) {
	var res []int

	for _, num := range arr {
		if num == 0 {
			res = append(res, 0)
			res = append(res, 0)
		} else {
			res = append(res, num)
		}
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = res[i]
	}
}
