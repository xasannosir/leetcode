package solution

func maxAscendingSum(nums []int) int {
	sum, temp, res := 0, 0, 0

	for _, num := range nums {
		if num > temp {
			sum += num
		} else {
			res = max(res, sum)
			sum = num
		}

		temp = num
	}

	res = max(res, sum)

	return res
}
