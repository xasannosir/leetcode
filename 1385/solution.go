package solution

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	total := 0

	abs := func(num int) int {
		if num < 0 {
			return -num
		}

		return num
	}

	for i := 0; i < len(arr1); i++ {
		isHave := true

		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				isHave = false
				break
			}
		}

		if isHave {
			total++
		}
	}

	return total
}
