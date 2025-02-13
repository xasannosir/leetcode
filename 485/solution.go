package solution

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	res := 0

	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			res = max(res, count)
			count = 0
		}
	}

	res = max(res, count)

	return res
}
