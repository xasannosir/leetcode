package solution

func targetIndices(nums []int, target int) []int {
	sort := func(arr []int) {
		for i := 0; i < len(arr); i++ {
			flag := false
			for j := 0; j < len(arr)-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					flag = true
				}
			}

			if !flag {
				break
			}
		}
	}

	sort(nums)

	res := make([]int, 0)

	for i, val := range nums {
		if val == target {
			res = append(res, i)
		}
	}

	return res
}
